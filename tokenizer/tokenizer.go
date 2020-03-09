package tokenizer

import (
	"fmt"
)

type Token struct {
	Typ   TokenType
	Value interface{}
}

type Tokenizer struct {
	tokens        []Token
	input         string
	currIdx       int
	currToken     int
	currLine      int
	currPosInLine int
}

func (t *Tokenizer) Print() {
	for _, t := range t.tokens {
		fmt.Printf("<%s:%s>\n", t.Value, t.Typ)
	}
}

func (t Tokenizer) currChar() byte {
	return t.input[t.currIdx]
}

func (t *Tokenizer) Tokenize() error {
	for t.currIdx < len(t.input) {
		token, ln, err := t.getTokenType()
		if err != nil {
			return fmt.Errorf("met unsupported token starting at line: %d, pos: %d (symbol:%c)",
				t.currLine+1, t.currPosInLine+1, t.currChar())
		}
		if token.Typ != Space {
			t.tokens = append(t.tokens, token)
		}
		t.currIdx += ln
		t.currPosInLine += ln
		if token.Typ == LineBreak {
			t.currLine++
			t.currPosInLine = 0
		}
	}
	return nil
}

func (t *Tokenizer) GetNextToken() Token {
	if t.currToken < len(t.tokens) {
		tok := t.tokens[t.currToken]
		t.currToken++
		return tok
	}
	return Token{Typ: EOF, Value: ""}

}

func NewTokenizer(inp string) Tokenizer {
	return Tokenizer{input: inp}
}
