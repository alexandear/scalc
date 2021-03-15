package scalc

import (
	"fmt"
)

type Expression struct {
	Operator Operator `parser:"\"[\" @Operator"`
	N        uint     `parser:"      @Positive"`
	Sets     []*Set   `parser:"      @@+ \"]\""`
}

type Operator string

const (
	OpEQ Operator = "EQ"
	OpLE Operator = "LE"
	OpGR Operator = "GR"
)

type Set struct {
	File          *string     `parser:"  @File"`
	SubExpression *Expression `parser:"| @@"`
}

func (s *Set) String() string {
	var res string

	if s.File != nil {
		res += *s.File
	} else if s.SubExpression != nil {
		res += fmt.Sprintf("%+v", s.SubExpression)
	}

	return res
}
