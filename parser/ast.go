package parser

import (
	"calculator_ast/tokenizer"
)

type ProgramNode struct {
	ChampName    string
	ChampComment string
	Code         CodeNode
}

type CodeNode struct {
	Commands []CommandNode
}

type CommandNode struct {
	Labels      []LabelNode
	Instruction *InstructionNode
}

type LabelNode struct {
	Token       tokenizer.Token
	Name        string
	OffsetBytes int
}

type InstructionNode struct {
	Token       tokenizer.Token
	Name        tokenizer.InstructionName
	Meta        tokenizer.InstructionMeta
	Args        []InstructionArgument
	OffsetBytes int
	Size        int // for debug, not used in logic
}

type InstructionArgument struct {
	Token tokenizer.Token
	Type  tokenizer.ArgumentType
	Value interface{}
}
