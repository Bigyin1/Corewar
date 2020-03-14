package compiler

import "calculator_ast/parser"

type InstructionOffsetIndexer struct {
	ast parser.ProgramNode
}

func (ip InstructionOffsetIndexer) evalInstrSize(instr *parser.InstructionNode) int {
	size := 1
	if instr.Meta.IsArgTypeCode {
		size++
	}
	for _, arg := range instr.Args {
		size += arg.Type.Size
	}
	return size
}

func (ip InstructionOffsetIndexer) SetOffsets() {
	currOffset := 0
	for _, cmd := range ip.ast.Code.Commands {
		if cmd.Instruction == nil {
			for idx := range cmd.Labels {
				cmd.Labels[idx].OffsetBytes = currOffset
			}
			continue
		}
		cmd.Instruction.OffsetBytes = currOffset
		for idx := range cmd.Labels {
			cmd.Labels[idx].OffsetBytes = currOffset
		}
		instrSize := ip.evalInstrSize(cmd.Instruction)
		cmd.Instruction.Size = instrSize
		currOffset += instrSize
	}
}

func NewInstructionOffsetIndexer(ast parser.ProgramNode) InstructionOffsetIndexer {
	return InstructionOffsetIndexer{ast: ast}
}
