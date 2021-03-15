package parser

import (
	"github.com/alexandear/scalc/pkg/scalc"
)

type Expression struct {
	Operator scalc.Operator `parser:"\"[\" @Operator"`
	N        uint           `parser:"      @Positive"`
	Sets     []*Set         `parser:"      @@+ \"]\""`
}

type Set struct {
	File          *string     `parser:"  @File"`
	SubExpression *Expression `parser:"| @@"`
}
