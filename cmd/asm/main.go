package main

import (
	"calculator_ast/pkg/asm/compiler"
	"calculator_ast/pkg/asm/parser"
	"calculator_ast/pkg/asm/tokenizer"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func compile(input string) (*compiler.Compiler, error) {
	lex := tokenizer.NewTokenizer(input)
	err := lex.Tokenize()
	if err != nil {
		return nil, err
	}

	pars := parser.NewParser(lex)
	ast, err := pars.Parse()
	if err != nil {
		return nil, err
	}
	comp := compiler.NewCompiler(ast)
	err = comp.Compile()
	if err != nil {
		return nil, err
	}
	return &comp, nil
}

func main() {
	var printDebug bool
	flag.BoolVar(&printDebug, "d", false, "print annotated result code")
	flag.Parse()
	if len(flag.Args()) == 1 {
		inputFile := flag.Arg(0)
		if !strings.HasSuffix(inputFile, ".asm") && !strings.HasSuffix(inputFile, ".s") {
			fmt.Println("accept files .s or .asm")
			return
		}
		input, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		comp, err := compile(string(input))
		if err != nil {
			fmt.Println(err)
			return
		}
		if printDebug {
			comp.PrintAnnotatedCode()
			return
		}
		outfile := ""
		if strings.HasSuffix(inputFile, ".asm") {
			outfile = strings.TrimSuffix(inputFile, ".asm")
		}
		if strings.HasSuffix(inputFile, ".s") {
			outfile = strings.TrimSuffix(inputFile, ".s")
		}
		outfile += ".cor"
		err = ioutil.WriteFile(outfile, comp.GetByteCode(), 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
	fmt.Println("Usage: ./asm [-d] <sourcefile.s>")
}
