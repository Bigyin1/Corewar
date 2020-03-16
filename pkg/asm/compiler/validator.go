package compiler

import (
	"calculator_ast/consts"
	"calculator_ast/pkg/asm/parser"
	"fmt"
)

func (c Compiler) getExpectedArgTypes(expArgCodes uint8) string {
	res := ""
	if expArgCodes&consts.TRegIdCode != 0 {
		res += consts.TReg.ArgTypeName
	}
	if expArgCodes&consts.TDirIdCode != 0 {
		if len(res) != 0 {
			res += " or "
		}
		res += consts.TDir.ArgTypeName
	}
	if expArgCodes&consts.TIndIdCode != 0 {
		if len(res) != 0 {
			res += " or "
		}
		res += consts.TInd.ArgTypeName
	}
	return res
}

func (c Compiler) validateInstruction(cmd *parser.InstructionNode) error {
	currCmdArgs := cmd.Args
	cmdExpArgs := cmd.Meta.AllowedArgs
	if len(currCmdArgs) != len(cmdExpArgs) {
		return fmt.Errorf("invalid args count for instruction %s", cmd.Name)

	}
	for idx, expArgCodes := range cmdExpArgs {
		currArgType := currCmdArgs[idx].Type
		if expArgCodes&currArgType.ArgTypeIDCode == 0 {
			return fmt.Errorf("invalid arg type %s for cmd %s, expected %s; line: %d, col:%d",
				currArgType.ArgTypeName, cmd.Name, c.getExpectedArgTypes(expArgCodes),
				cmd.Token.PosLine, cmd.Token.PosColumn)
		}
	}
	return nil
}

func (c Compiler) validateInstructions() error {
	for _, cmd := range c.ast.Code.Commands {
		if cmd.Instruction != nil {
			err := c.validateInstruction(cmd.Instruction)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
