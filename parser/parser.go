package parser

import (
	"calculator_ast/tokenizer"
	"fmt"
)

type Parser struct {
	lexer        tokenizer.Tokenizer
	currentToken tokenizer.Token
}

func (p *Parser) eatToken(tokenType tokenizer.TokenType) error {
	if p.currentToken.Typ == tokenType {
		p.currentToken = p.lexer.GetNextToken()
		return nil
	}
	return fmt.Errorf("unexpected token: %s, exp: %s", tokenType, p.currentToken)
}

func (p *Parser) factor() (AST, error) {
	if p.currentToken.Typ == tokenizer.Integer {
		token := p.currentToken
		err := p.eatToken(p.currentToken.Typ)
		if err != nil {
			return nil, err
		}
		return &NumASTNode{token: token}, nil
	} else if p.currentToken.Typ == tokenizer.Lparen {
		err := p.eatToken(tokenizer.Lparen)
		if err != nil {
			return nil, err
		}
		expr, err := p.expr()
		if err != nil {
			return nil, err
		}
		err = p.eatToken(tokenizer.Rparen)
		if err != nil {
			return nil, err
		}
		return expr, nil
	}
	return nil, fmt.Errorf("got %v, expected number or expr in brackets", p.currentToken)
}

func (p *Parser) term() (AST, error) {
	node, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.currentToken.Typ == tokenizer.Mult || p.currentToken.Typ == tokenizer.Div {
		token := p.currentToken
		err := p.eatToken(p.currentToken.Typ)
		if err != nil {
			return nil, err
		}
		rf, err := p.factor()
		if err != nil {
			return nil, err
		}
		node = &BinOpASTNode{left: node, token: token, right: rf}
	}
	return node, nil
}

func (p *Parser) expr() (AST, error) {
	node, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.currentToken.Typ == tokenizer.Sum || p.currentToken.Typ == tokenizer.Sub {
		token := p.currentToken
		err := p.eatToken(p.currentToken.Typ)
		if err != nil {
			return nil, err
		}
		rt, err := p.term()
		if err != nil {
			return nil, err
		}
		node = &BinOpASTNode{left: node, token: token, right: rt}
	}
	return node, nil

}

func (p *Parser) Parse() (AST, error) {
	p.currentToken = p.lexer.GetNextToken()
	ast, err := p.expr()
	if err != nil {
		return nil, err
	}
	return ast, nil
}

func NewParser(lex tokenizer.Tokenizer) Parser {
	return Parser{lexer: lex}
}
