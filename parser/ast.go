package parser

import "calculator_ast/tokenizer"

type AST interface {
	GetChildren() []AST
	GetToken() tokenizer.Token
}
type BinOpASTNode struct {
	left  AST
	right AST
	token tokenizer.Token
}

func (b *BinOpASTNode) GetChildren() []AST {
	return []AST{b.left, b.right}
}

func (b *BinOpASTNode) GetToken() tokenizer.Token {
	return b.token
}

type NumASTNode struct {
	token tokenizer.Token
}

func (b *NumASTNode) GetChildren() []AST {
	return []AST{}
}

func (b *NumASTNode) GetToken() tokenizer.Token {
	return b.token
}
