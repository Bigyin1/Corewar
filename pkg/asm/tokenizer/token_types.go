package tokenizer

import (
	"errors"

	"github.com/Bigyin1/Corewar/pkg/consts"
)

type ArgumentType struct {
	ArgTypeIDCode consts.TypeID
	ArgTypeName   string
	ByteCode      uint8
	Size          int
}

var (
	TReg = ArgumentType{consts.TRegIdCode, "T_REG",
		consts.TypeIDToByteCode(consts.TRegIdCode),
		consts.RegArgSize}
	TDir = ArgumentType{consts.TDirIdCode, "T_DIR",
		consts.TypeIDToByteCode(consts.TDirIdCode),
		consts.DirArgSize}
	TInd = ArgumentType{consts.TIndIdCode, "T_IND",
		consts.TypeIDToByteCode(consts.TIndIdCode),
		consts.IndArgSize}
)

type TokenType string

func (tt TokenType) IsOfArgType() bool {
	if tt == Direct || tt == DirectLabel || tt == Indirect || tt == IndirectLabel || tt == Register {
		return true
	}
	return false
}

func (tt TokenType) IsDirectArgType() bool {
	if tt == Direct || tt == DirectLabel {
		return true
	}
	return false
}

func (tt TokenType) IsIndirectArgType() bool {
	if tt == Indirect || tt == IndirectLabel {
		return true
	}
	return false
}

func (tt TokenType) IsRegisterArgType() bool {
	if tt == Register {
		return true
	}
	return false
}

func (tt TokenType) GetArgType() ArgumentType {
	if tt == Register {
		return TReg
	}
	if tt == Direct {
		return TDir
	}
	if tt == DirectLabel {
		return TDir
	}
	if tt == Indirect {
		return TInd
	}
	if tt == IndirectLabel {
		return TInd
	}
	return ArgumentType{}
}

var EOFErr = errors.New("EOF")

const (
	Str           TokenType = "STRING"
	ChampName               = "CHAMP_NAME"
	ChampComment            = "CHAMP_COMMENT"
	Instr                   = "INSTRUCTION"
	Space                   = "SPACE"
	Label                   = "LABEL"
	Separator               = "SEPARATOR"
	Register                = "REGISTER"
	Direct                  = "DIRECT"
	DirectLabel             = "DIRECT_LABEL"
	Indirect                = "INDIRECT"
	IndirectLabel           = "INDIRECT_LABEL"
	LineBreak               = "LINE_BREAK"
	Sum                     = "SUM"
	Sub                     = "SUB"
	Comment                 = "COMMENT"
	EOF                     = "EOF"
)
