package main

import (
	"calculator_ast/compiler"
	"calculator_ast/parser"
	"calculator_ast/tokenizer"
	"fmt"
	"io/ioutil"
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

	pars := parser.NewParser(lex)
	res, err := pars.Parse()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	instrValidator := compiler.NewInstructionsValidator(res)
	instrValidator.ValidateInstructions()
	return ""
}

func main() {
	if len(os.Args) == 2 {
		file := os.Args[1]
		input, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		res := evaluate(string(input))
		fmt.Println(res)
		return
	}
	file := ""
	_, _ = fmt.Fscanf(os.Stdin, "%s", &file)
	input, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	res := evaluate(string(input))
	fmt.Println(res)
}
