package consts

const NameHeader = ".name"
const ChampNameLength = 64
const CommentHeader = ".comment"
const LabelChars = "abcdefghijklmnopqrstuvwxyz_0123456789"
const RegNumber = 16
const SeparatorSymbol = ","

const RegSize = 1
const ShortDirSize = 2
const IndSize = 2
const DirSize = 4

//type RegType uint8
//type IndType int16
//type ShortDirType int16
//type DirType int32

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
	ArgTypeIDCode uint8
	ArgTypeName   string
	ByteCode      uint8
	Size          int
}

// Type codes for internal usage, not for byte code
const (
	TRegIdCode = 0b001
	TDirIdCode = 0b010
	TIndIdCode = 0b100
)

var (
	TDir = ArgumentType{TDirIdCode, "T_DIR", 0b10, DirSize}
	TReg = ArgumentType{TRegIdCode, "T_REG", 0b01, RegSize}
	TInd = ArgumentType{TIndIdCode, "T_IND", 0b11, IndSize}
)

type InstructionMeta struct {
	AllowedArgs   []uint8
	TDirSize      int
	OpCode        byte
	IsArgTypeCode bool
}

var InstructionsConfig = map[InstructionName]InstructionMeta{
	LIVE: {AllowedArgs: []uint8{TDirIdCode}, IsArgTypeCode: false, TDirSize: 4, OpCode: 0x01},
	LD:   {AllowedArgs: []uint8{TDirIdCode | TIndIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: 4, OpCode: 0x02},
	ST:   {AllowedArgs: []uint8{TRegIdCode, TRegIdCode | TIndIdCode}, IsArgTypeCode: true, TDirSize: 4, OpCode: 0x03},
	ADD:  {AllowedArgs: []uint8{TRegIdCode, TRegIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: 4, OpCode: 0x04},
	SUB:  {AllowedArgs: []uint8{TRegIdCode, TRegIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: 4, OpCode: 0x05},
	AND: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: 4, OpCode: 0x06},
	OR: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: 4, OpCode: 0x07},
	XOR: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: 4, OpCode: 0x08},
	ZJMP: {AllowedArgs: []uint8{TDirIdCode}, IsArgTypeCode: false, TDirSize: 2, OpCode: 0x09},
	LDI: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: 2, OpCode: 0x0A},
	STI: {AllowedArgs: []uint8{TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode},
		IsArgTypeCode: true, TDirSize: 2, OpCode: 0x0B},
	FORK: {AllowedArgs: []uint8{TDirIdCode}, IsArgTypeCode: false, TDirSize: 2, OpCode: 0x0C},
	LLD:  {AllowedArgs: []uint8{TDirIdCode | TIndIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: 4, OpCode: 0x0D},
	LLDI: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: 2, OpCode: 0x0E},
	LFORK: {AllowedArgs: []uint8{TDirIdCode}, IsArgTypeCode: false, TDirSize: 2, OpCode: 0x0F},
	AFF:   {AllowedArgs: []uint8{TRegIdCode}, IsArgTypeCode: false, TDirSize: 4, OpCode: 0x10},
}
