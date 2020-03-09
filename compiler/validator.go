package compiler

import (
	"calculator_ast/parser"
	"calculator_ast/tokenizer"
	"fmt"
)

type InstructionsValidator struct {
	ast parser.ProgramNode
}

func (v InstructionsValidator) getExpectedArgTypes(expArgCodes uint8) string {
	res := ""
	if expArgCodes&tokenizer.T_REG_CODE != 0 {
		res += tokenizer.T_REG.ArgTypeName
	}
	if expArgCodes&tokenizer.T_DIR_CODE != 0 {
		if len(res) != 0 {
			res += " or "
		}
		res += tokenizer.T_DIR.ArgTypeName
	}
	if expArgCodes&tokenizer.T_IND_CODE != 0 {
		if len(res) != 0 {
			res += " or "
		}
		res += tokenizer.T_IND.ArgTypeName
	}
	return res
}

func (v InstructionsValidator) validateInstruction(cmd *parser.InstructionNode) bool {
	currCmdArgs := cmd.Args
	cmdExpArgs := tokenizer.Instructions[cmd.InstructionName]
	if len(currCmdArgs) != len(cmdExpArgs) {
		fmt.Printf("invalid args count for instruction %s", cmd.InstructionName)
		return false
	}
	for idx, expArgCodes := range cmdExpArgs {
		currArgType := currCmdArgs[idx].Type
		if expArgCodes&currArgType.ArgTypeCode == 0 {
			fmt.Printf("invalid arg type %s for cmd %s, expected %s",
				currArgType.ArgTypeName, cmd.InstructionName, v.getExpectedArgTypes(expArgCodes))
			return false
		}
	}
	return true
}

func (v InstructionsValidator) ValidateInstructions() bool {
	for _, cmd := range v.ast.Code.Commands {
		isValid := v.validateInstruction(cmd.Instruction)
		if !isValid {
			return false
		}
	}
	return true
}

func NewInstructionsValidator(ast parser.ProgramNode) InstructionsValidator {
	return InstructionsValidator{ast: ast}
}
