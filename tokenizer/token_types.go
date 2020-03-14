package tokenizer

import "calculator_ast/consts"

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
		return consts.T_REG
	}
	if tt == Direct {
		return consts.T_DIR
	}
	if tt == DirectLabel {
		return consts.T_DIR
	}
	if tt == Indirect {
		return consts.T_IND
	}
	if tt == IndirectLabel {
		return consts.T_IND
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
