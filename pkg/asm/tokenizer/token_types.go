package tokenizer

import "corewar/consts"

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

func (tt TokenType) GetArgType() consts.ArgumentType {
	if tt == Register {
		return consts.TReg
	}
	if tt == Direct {
		return consts.TDir
	}
	if tt == DirectLabel {
		return consts.TDir
	}
	if tt == Indirect {
		return consts.TInd
	}
	if tt == IndirectLabel {
		return consts.TInd
	}
	return consts.ArgumentType{}
}

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
	EOF                     = "EOF"
)
