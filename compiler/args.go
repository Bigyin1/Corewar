package compiler

import (
	"calculator_ast/consts"
	"calculator_ast/tokenizer"
	"fmt"
)

func (c Compiler) SetupLabelsTable() error {
	for _, cmd := range c.ast.Code.Commands {
		for _, label := range cmd.Labels {
			_, ok := c.labelTable[label.Name]
			if ok {
				return fmt.Errorf("duplicate label %s at line: %d, col: %d", label.Name,
					label.Token.PosLine, label.Token.PosColumn)
			}
			c.labelTable[label.Name] = label.OffsetBytes
		}
	}
	return nil
}

func (c Compiler) FillArgValues() error {
	for _, cmd := range c.ast.Code.Commands {
		if cmd.Instruction == nil {
			continue
		}
		cmdOffset := cmd.Instruction.OffsetBytes
		for idx := range cmd.Instruction.Args {
			arg := &cmd.Instruction.Args[idx]
			switch arg.Token.Type {
			case tokenizer.Register:
				arg.Value = arg.Token.Value.(tokenizer.RegisterTokenVal).GetValue()
			case tokenizer.Direct:
				val := arg.Token.Value.(tokenizer.DirectTokenVal).GetValue()
				arg.Value = val
				if arg.Type.Size == consts.SHORT_DIR_SIZE {
					arg.Value = int16(val)
				}
			case tokenizer.DirectLabel:
				label := arg.Token.Value.(tokenizer.DirectLabelTokenVal).GetValue()
				labelOffset, ok := c.labelTable[label]
				if !ok {
					return fmt.Errorf("failed to find %s label, line: %d", label, cmd.Instruction.Token.PosLine)
				}
				arg.Value = labelOffset - cmdOffset
				if arg.Type.Size == consts.SHORT_DIR_SIZE {
					arg.Value = int16(labelOffset - cmdOffset)
				}
			case tokenizer.Indirect:
				arg.Value = arg.Token.Value.(tokenizer.IndirectTokenVal).GetValue()
			case tokenizer.IndirectLabel:
				label := arg.Token.Value.(tokenizer.IndirectLabelTokenVal).GetValue()
				labelOffset, ok := c.labelTable[label]
				if !ok {
					return fmt.Errorf("failed to find %s label, line: %d", label, cmd.Instruction.Token.PosLine)
				}
				arg.Value = int16(labelOffset - cmdOffset)
			}

		}
	}
	return nil
}
