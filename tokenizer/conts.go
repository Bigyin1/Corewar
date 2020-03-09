package tokenizer

const nameHeader = ".name"
const commentHeader = ".comment"
const labelChars = "abcdefghijklmnopqrstuvwxyz_0123456789"
const RegNumber = 16
const SeparatorSymbol = ","

type InstructionName string

const (
	LIVE  InstructionName = "live"
	LD                    = "ld"
	ST                    = "st"
	ADD                   = "add"
	SUB                   = "sub"
	AND                   = "and"
	OR                    = "or"
	XOR                   = "xor"
	ZJMP                  = "zjmp"
	LDI                   = "ldi"
	STI                   = "sti"
	FORK                  = "fork"
	LLD                   = "lld"
	LLDI                  = "lldi"
	LFORK                 = "lfork"
	AFF                   = "aff"
)

type ArgumentType struct {
	ArgTypeCode uint8
	ArgTypeName string
}

const (
	T_REG_CODE = 0b001
	T_DIR_CODE = 0b010
	T_IND_CODE = 0b100
)

var (
	T_DIR = ArgumentType{T_DIR_CODE, "T_DIR"}
	T_REG = ArgumentType{T_REG_CODE, "T_REG"}
	T_IND = ArgumentType{T_IND_CODE, "T_IND"}
)

var Instructions = map[InstructionName][]uint8{
	LIVE:  {T_DIR_CODE},
	LD:    {T_DIR_CODE | T_IND_CODE, T_REG_CODE},
	ST:    {T_REG_CODE, T_REG_CODE | T_IND_CODE},
	ADD:   {T_REG_CODE, T_REG_CODE, T_REG_CODE},
	SUB:   {T_REG_CODE, T_REG_CODE, T_REG_CODE},
	AND:   {T_DIR_CODE | T_IND_CODE | T_REG_CODE, T_DIR_CODE | T_IND_CODE | T_REG_CODE, T_REG_CODE},
	OR:    {T_DIR_CODE | T_IND_CODE | T_REG_CODE, T_DIR_CODE | T_IND_CODE | T_REG_CODE, T_REG_CODE},
	XOR:   {T_DIR_CODE | T_IND_CODE | T_REG_CODE, T_DIR_CODE | T_IND_CODE | T_REG_CODE, T_REG_CODE},
	ZJMP:  {T_DIR_CODE},
	LDI:   {T_DIR_CODE | T_IND_CODE | T_REG_CODE, T_REG_CODE | T_DIR_CODE, T_REG_CODE},
	STI:   {T_REG_CODE, T_DIR_CODE | T_IND_CODE | T_REG_CODE, T_REG_CODE | T_DIR_CODE},
	FORK:  {T_DIR_CODE},
	LLD:   {T_DIR_CODE | T_IND_CODE, T_REG_CODE},
	LLDI:  {T_DIR_CODE | T_IND_CODE | T_REG_CODE, T_REG_CODE | T_DIR_CODE, T_REG_CODE},
	LFORK: {T_DIR_CODE},
	AFF:   {T_REG_CODE},
}
