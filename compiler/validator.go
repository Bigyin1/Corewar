package compiler

import (
	"calculator_ast/consts"
	"calculator_ast/parser"
	"fmt"
)

type InstructionsValidator struct {
	ast parser.ProgramNode
}

func (v InstructionsValidator) getExpectedArgTypes(expArgCodes uint8) string {
	res := ""
	if expArgCodes&consts.T_REG_ID_CODE != 0 {
		res += consts.T_REG.ArgTypeName
	}
	if expArgCodes&consts.T_DIR_ID_CODE != 0 {
		if len(res) != 0 {
			res += " or "
		}
		res += consts.T_DIR.ArgTypeName
	}
	if expArgCodes&consts.T_IND_ID_CODE != 0 {
		if len(res) != 0 {
			res += " or "
		}
		res += consts.T_IND.ArgTypeName
	}
	return res
}

func (v InstructionsValidator) validateInstruction(cmd *parser.InstructionNode) bool {
	currCmdArgs := cmd.Args
	cmdExpArgs := cmd.Meta.AllowedArgs
	if len(currCmdArgs) != len(cmdExpArgs) {
		fmt.Printf("invalid args count for instruction %s", cmd.Name)
		return false
	}
	for idx, expArgCodes := range cmdExpArgs {
		currArgType := currCmdArgs[idx].Type
		if expArgCodes&currArgType.ArgTypeIDCode == 0 {
			fmt.Printf("invalid arg type %s for cmd %s, expected %s; line: %d, col:%d",
				currArgType.ArgTypeName, cmd.Name, v.getExpectedArgTypes(expArgCodes),
				cmd.Token.PosLine, cmd.Token.PosColumn)
			return false
		}
	}
	return true
}

func (v InstructionsValidator) ValidateInstructions() bool {
	for _, cmd := range v.ast.Code.Commands {
		if cmd.Instruction != nil {
			isValid := v.validateInstruction(cmd.Instruction)
			if !isValid {
				return false
			}
		}
	}
	return true
}

func NewInstructionsValidator(ast parser.ProgramNode) InstructionsValidator {
	return InstructionsValidator{ast: ast}
}
