package tokenizer

import (
	"fmt"
)

type Token struct {
	Type      TokenType
	Value     interface{}
	PosLine   int
	PosColumn int
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
		fmt.Printf("<%s:%s>\n", t.Value, t.Type)
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
				t.currLine, t.currPosInLine, t.currChar())
		}
		if token.Type != Space {
			token.PosLine = t.currLine
			token.PosColumn = t.currPosInLine
			t.tokens = append(t.tokens, token)
		}
		t.currIdx += ln
		t.currPosInLine += ln
		if token.Type == LineBreak {
			t.currLine++
			t.currPosInLine = 1
		}
	}
	t.tokens = append(t.tokens)
	return nil
}

func (t *Tokenizer) GetNextToken() Token {
	if t.currToken < len(t.tokens) {
		tok := t.tokens[t.currToken]
		t.currToken++
		return tok
	}
	return Token{Type: EOF, Value: "", PosLine: t.tokens[len(t.tokens)-1].PosLine + 1}

}

func NewTokenizer(inp string) Tokenizer {
	return Tokenizer{input: inp, currLine: 1}
}
