package main

import (
	"calculator_ast/interpreter"
	"calculator_ast/parser"
	"calculator_ast/tokenizer"
	"fmt"
	"os"
)

func evaluate(input string) string {
	lex := tokenizer.NewTokenizer(input)
	err := lex.Tokenize()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	lex.Print()
	prs := parser.NewParser(lex)
	ast, err := prs.Parse()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	intr := interpreter.NewInterpreter(ast)
	result := intr.Evaluate()
	return fmt.Sprint(result)
}

func main() {
	input := ""
	if len(os.Args) == 2 {
		input = os.Args[1]
		res := evaluate(input)
		fmt.Println(res)
		return
	}
	_, _ = fmt.Fscanf(os.Stdin, "%s", &input)
	res := evaluate(input)
	fmt.Println(res)
}
