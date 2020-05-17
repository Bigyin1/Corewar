package compiler

import "corewar/pkg/asm/parser"

func (c Compiler) evalInstrSize(instr *parser.InstructionNode) int {
	size := 1
	if instr.Meta.IsArgTypeCode {
		size++
	}
	for _, arg := range instr.Args {
		size += arg.Type.Size
	}
	return size
}

func (c Compiler) setOffsets() {
	currOffset := 0
	for _, cmd := range c.ast.Code.Commands {
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
		instrSize := c.evalInstrSize(cmd.Instruction)
		cmd.Instruction.Size = instrSize
		currOffset += instrSize
	}
}
