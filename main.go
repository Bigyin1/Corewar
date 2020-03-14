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
	off := compiler.NewInstructionOffsetIndexer(res)
	off.SetOffsets()
	c := compiler.NewCompiler(res)
	err = c.SetupLabelsTable()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	err = c.FillArgValues()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//code := c.GetByteCode()
	c.PrintAnnotatedCode()
	//for _, b := range code {
	//	fmt.Printf("%02x ", b)
	//}
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
