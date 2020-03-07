package tokenizer

import (
	"fmt"
	"strconv"
	"strings"
)

type TokenType string

const (
	Sum     TokenType = "SUM"
	Sub               = "SUB"
	Mult              = "MULT"
	Div               = "DIV"
	Lparen            = "LPAREN"
	Rparen            = "RPAREN"
	Integer           = "INTEGER"
	EOF               = "EOF"
)

type Token struct {
	Typ   TokenType
	Value interface{}
}

type Tokenizer struct {
	tokens    []Token
	input     string
	currIdx   int
	currToken int
}

func (t *Tokenizer) Print() {
	for _, t := range t.tokens {
		if t.Typ != Integer {
			fmt.Printf("<%c:%s>\n", t.Value, t.Typ)
			continue
		}
		fmt.Printf("<%v:%s>\n", t.Value, t.Typ)
	}
}

func (t *Tokenizer) parserInt() int {
	res := strings.Builder{}
	for t.currIdx < len(t.input) {
		if t.currChar() >= '0' && t.currChar() <= '9' {
			res.WriteByte(t.currChar())
			t.currIdx++
			continue
		}
		break
	}
	i, _ := strconv.Atoi(res.String())
	return i
}

func (t Tokenizer) currChar() byte {
	return t.input[t.currIdx]
}

func (t *Tokenizer) skipSpaces() {
	for t.currIdx < len(t.input) {
		if t.currChar() == ' ' {
			t.currIdx++
			continue
		}
		return
	}
}

func (t *Tokenizer) Tokenize() error {
	for t.currIdx < len(t.input) {
		if t.currChar() == ' ' {
			t.skipSpaces()
			continue
		}
		if t.currChar() == '+' {
			t.tokens = append(t.tokens, Token{Typ: Sum, Value: '+'})
			t.currIdx++
			continue
		}
		if t.currChar() == '-' {
			t.tokens = append(t.tokens, Token{Typ: Sub, Value: '-'})
			t.currIdx++
			continue
		}
		if t.currChar() == '*' {
			t.tokens = append(t.tokens, Token{Typ: Mult, Value: '*'})
			t.currIdx++
			continue
		}
		if t.currChar() == '/' {
			t.tokens = append(t.tokens, Token{Typ: Div, Value: '/'})
			t.currIdx++
			continue
		}
		if t.currChar() == '(' {
			t.tokens = append(t.tokens, Token{Typ: Lparen, Value: '('})
			t.currIdx++
			continue
		}
		if t.currChar() == ')' {
			t.tokens = append(t.tokens, Token{Typ: Rparen, Value: ')'})
			t.currIdx++
			continue
		}
		if t.currChar() >= '0' && t.currChar() <= '9' {
			t.tokens = append(t.tokens, Token{Typ: Integer, Value: t.parserInt()})
			continue
		}
		return fmt.Errorf("met unsupported symbol: %c", t.currChar())
	}
	return nil
}

func (t *Tokenizer) GetNextToken() Token {
	if t.currToken < len(t.tokens) {
		tok := t.tokens[t.currToken]
		t.currToken++
		return tok
	}
	return Token{Typ: EOF, Value: nil}

}

func NewTokenizer(inp string) Tokenizer {
	return Tokenizer{input: inp}
}
