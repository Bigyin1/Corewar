package interpreter

import (
	"calculator_ast/parser"
	"calculator_ast/tokenizer"
)

type Interpreter struct {
	ast parser.AST
}

func (i *Interpreter) visit(node parser.AST) int {
	res := 0
	switch node.(type) {
	case *parser.BinOpASTNode:
		res = i.visitBinOpNode(node)
	case *parser.NumASTNode:
		res = i.visitNumNode(node)
	}
	return res
}

func (i *Interpreter) visitBinOpNode(node parser.AST) int {
	lval := i.visit(node.GetChildren()[0])
	rval := i.visit(node.GetChildren()[1])
	switch node.GetToken().Typ {
	case tokenizer.Sum:
		return lval + rval
	case tokenizer.Sub:
		return lval - rval
	case tokenizer.Mult:
		return lval * rval
	case tokenizer.Div:
		return lval / rval
	}
	return lval + rval
}

func (i *Interpreter) visitNumNode(node parser.AST) int {
	return node.GetToken().Value.(int)
}

func (i *Interpreter) Evaluate() int {
	return i.visit(i.ast)
}

func NewInterpreter(ast parser.AST) Interpreter {
	return Interpreter{ast: ast}
}
