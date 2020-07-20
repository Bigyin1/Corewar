package parser

import (
	"corewar/pkg/asm/tokenizer"
	"corewar/pkg/consts"
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
	Name        consts.InstructionName
	Meta        consts.InstructionMeta
	Args        []InstructionArgument
	OffsetBytes int
	Size        int // for debug, not used in logic
}

type InstructionArgument struct {
	Token tokenizer.Token
	Type  consts.ArgumentType
	Value interface{}
}
