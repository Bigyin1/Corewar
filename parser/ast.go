package parser

import (
	"calculator_ast/tokenizer"
	"fmt"
)

type ProgramNode struct {
	ChampName    string
	ChampComment string
	Code         CodeNode
}

func (pr ProgramNode) PrintTree() {
	fmt.Print("\t\t")
	fmt.Print(pr.ChampName)
	fmt.Print("\t")
	fmt.Print(pr.ChampComment)

}

type CodeNode struct {
	Commands []CommandNode
}

type CommandNode struct {
	Label       []LabelNode
	Instruction *InstructionNode
}

type LabelNode struct {
	Name string
}

type InstructionNode struct {
	InstructionName tokenizer.InstructionName
	Args            []InstructionArgument
}

type InstructionArgument struct {
	Token tokenizer.Token
	Type  tokenizer.ArgumentType
}
