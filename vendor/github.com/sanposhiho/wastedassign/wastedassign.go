package wastedassign

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/ssa"
)

const doc = "wastedassign finds wasted assignment statements."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "wastedassign",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

type wastedAssignStruct struct {
	pos    token.Pos
	reason string
}

func run(pass *analysis.Pass) (interface{}, error) {
	// Plundered from buildssa.Run.
	mode := ssa.NaiveForm
	prog := ssa.NewProgram(pass.Fset, mode)

	// Create SSA packages for all imports.
	// Order is not significant.
	created := make(map[*types.Package]bool)
	var createAll func(pkgs []*types.Package)
	createAll = func(pkgs []*types.Package) {
		for _, p := range pkgs {
			if !created[p] {
				created[p] = true
				prog.CreatePackage(p, nil, nil, true)
				createAll(p.Imports())
			}
		}
	}
	createAll(pass.Pkg.Imports())

	// Create and build the primary package.
	ssapkg := prog.CreatePackage(pass.Pkg, pass.Files, pass.TypesInfo, false)
	ssapkg.Build()

	var srcFuncs []*ssa.Function
	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			if fdecl, ok := decl.(*ast.FuncDecl); ok {

				// SSA will not build a Function
				// for a FuncDecl named blank.
				// That's arguably too strict but
				// relaxing it would break uniqueness of
				// names of package members.
				if fdecl.Name.Name == "_" {
					continue
				}

				// (init functions have distinct Func
				// objects named "init" and distinct
				// ssa.Functions named "init#1", ...)

				fn := pass.TypesInfo.Defs[fdecl.Name].(*types.Func)
				if fn == nil {
					return nil, fmt.Errorf("failed to get func's typesinfo")
				}

				f := ssapkg.Prog.FuncValue(fn)
				if f == nil {
					return nil, fmt.Errorf("failed to get func's SSA-form intermediate representation")
				}

				var addAnons func(f *ssa.Function)
				addAnons = func(f *ssa.Function) {
					srcFuncs = append(srcFuncs, f)
					for _, anon := range f.AnonFuncs {
						addAnons(anon)
					}
				}
				addAnons(f)
			}
		}
	}

	typeSwitchPos := map[int]bool{}
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder([]ast.Node{new(ast.TypeSwitchStmt)}, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.TypeSwitchStmt:
			typeSwitchPos[pass.Fset.Position(n.Pos()).Line] = true
		}
	})

	wastedAssignMap := []wastedAssignStruct{}

	for _, sf := range srcFuncs {
		for _, bl := range sf.Blocks {
			blCopy := *bl
			for _, ist := range bl.Instrs {
				blCopy.Instrs = rmInstrFromInstrs(blCopy.Instrs, ist)
				switch ist.(type) {
				case *ssa.Store:
					var buf [10]*ssa.Value
					for _, op := range ist.Operands(buf[:0]) {
						if (*op) != nil && opInLocals(sf.Locals, op) {
							if reason := isNextOperationToOpIsStore([]*ssa.BasicBlock{&blCopy}, op, nil); reason != notWasted {
								if ist.Pos() != 0 && !typeSwitchPos[pass.Fset.Position(ist.Pos()).Line] {
									wastedAssignMap = append(wastedAssignMap, wastedAssignStruct{
										pos:    ist.Pos(),
										reason: reason.String(),
									})
								}
							}
						}
					}
				}
			}
		}
	}
	for _, was := range wastedAssignMap {
		pass.Reportf(was.pos, was.reason)
	}
	return nil, nil
}

type wastedReason string

const (
	noUseUntilReturn wastedReason = "reassigned, but never used afterwards"
	reassignedSoon   wastedReason = "wasted assignment"
	notWasted        wastedReason = ""
)

func (wr wastedReason) String() string {
	switch wr {
	case noUseUntilReturn:
		return "reassigned, but never used afterwards"
	case reassignedSoon:
		return "wasted assignment"
	case notWasted:
		return ""
	}
	return ""
}

func isNextOperationToOpIsStore(bls []*ssa.BasicBlock, currentOp *ssa.Value, haveCheckedMap *map[int]bool) wastedReason {
	wastedReasons := []wastedReason{}
	wastedReasonsCurrentBls := []wastedReason{}

	if haveCheckedMap == nil {
		haveCheckedMap = &map[int]bool{}
	}

	for _, bl := range bls {
		if (*haveCheckedMap)[bl.Index] == true {
			continue
		}
		(*haveCheckedMap)[bl.Index] = true
		breakFlag := false
		for _, ist := range bl.Instrs {
			if breakFlag {
				break
			}
			switch w := ist.(type) {
			case *ssa.Store:
				var buf [10]*ssa.Value
				for _, op := range ist.Operands(buf[:0]) {
					if *op == *currentOp {
						if w.Addr.Name() == (*currentOp).Name() {
							wastedReasonsCurrentBls = append(wastedReasonsCurrentBls, reassignedSoon)
							breakFlag = true
							break
						} else {
							return notWasted
						}
					}
				}
			default:
				var buf [10]*ssa.Value
				for _, op := range ist.Operands(buf[:0]) {
					if *op == *currentOp {
						// ??????store??????????????????
						return notWasted
					}
				}
			}
		}
		if len(bl.Succs) != 0 && !breakFlag {
			wastedReason := isNextOperationToOpIsStore(rmSameBlock(bl.Succs, bl), currentOp, haveCheckedMap)
			if wastedReason == notWasted {
				return notWasted
			}
			wastedReasons = append(wastedReasons, wastedReason)
		}
	}

	wastedReasons = append(wastedReasons, wastedReasonsCurrentBls...)

	if len(wastedReasons) != 0 {
		if containReassignedSoon(wastedReasons) {
			return reassignedSoon
		}
		return noUseUntilReturn
	}
	return noUseUntilReturn
}

func rmSameBlock(bls []*ssa.BasicBlock, currentBl *ssa.BasicBlock) []*ssa.BasicBlock {
	rto := []*ssa.BasicBlock{}

	for _, bl := range bls {
		if bl != currentBl {
			rto = append(rto, bl)
		}
	}
	return rto
}

func containNotWasted(ws []wastedReason) bool {
	for _, w := range ws {
		if w == notWasted {
			return true
		}
	}
	return false
}

func containReassignedSoon(ws []wastedReason) bool {
	for _, w := range ws {
		if w == reassignedSoon {
			return true
		}
	}
	return false
}

func rmInstrFromInstrs(instrs []ssa.Instruction, instrToRm ssa.Instruction) []ssa.Instruction {
	var rto []ssa.Instruction
	for _, i := range instrs {
		if i != instrToRm {
			rto = append(rto, i)
		}
	}
	return rto
}

func opInLocals(locals []*ssa.Alloc, op *ssa.Value) bool {
	for _, l := range locals {
		if *op == ssa.Value(l) {
			return true
		}
	}
	return false
}
