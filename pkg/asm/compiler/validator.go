package compiler

import (
	"fmt"

	"github.com/Bigyin1/Corewar/pkg/asm/parser"
	"github.com/Bigyin1/Corewar/pkg/asm/tokenizer"
	"github.com/Bigyin1/Corewar/pkg/consts"
)

func (c Compiler) getExpectedArgTypes(expArgCodes consts.TypeID) string {
	res := ""
	if expArgCodes&consts.TRegIdCode != 0 {
		res += tokenizer.TReg.ArgTypeName
	}
	if expArgCodes&consts.TDirIdCode != 0 {
		if len(res) != 0 {
			res += " or "
		}
		res += tokenizer.TDir.ArgTypeName
	}
	if expArgCodes&consts.TIndIdCode != 0 {
		if len(res) != 0 {
			res += " or "
		}
		res += tokenizer.TInd.ArgTypeName
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

func (c Compiler) validateMeta() error {
	if len(c.ast.ChampName) >= consts.ChampNameLength {
		return fmt.Errorf("champion name length must be less than %d", consts.ChampNameLength)
	}
	return nil
}
