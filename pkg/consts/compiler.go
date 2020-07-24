package consts

const NameHeader = ".name"
const ChampNameLength = 64
const CommentHeader = ".comment"
const LabelChars = "abcdefghijklmnopqrstuvwxyz_0123456789"
const RegNumber = 16
const SeparatorSymbol = ","

const MagicHeader = "\x00\xea\x83\xf3"
const NullSeq = "\x00\x00\x00\x00"

const ProgNameLength = 128
const CommentLength = 2048

const RegSize = 4

const RegArgSize = 1
const ShortDirSize = 2
const IndSize = 2
const DirSize = 4

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

type TypeID uint8

// Type codes for internal usage, not for byte code
const (
	TRegIdCode = 1 << 0
	TDirIdCode = 1 << 1
	TIndIdCode = 1 << 2
)

var (
	TDir = ArgumentType{TDirIdCode, "T_DIR", 0b10, DirSize}
	TReg = ArgumentType{TRegIdCode, "T_REG", 0b01, RegArgSize}
	TInd = ArgumentType{TIndIdCode, "T_IND", 0b11, IndSize}
)

type InstructionMeta struct {
	AllowedArgs   []uint8
	IsArgTypeCode bool
	TDirSize      int
	OpCode        byte
	CyclesToExec  int
	IdxMod        bool
}

var InstructionsConfig = map[InstructionName]InstructionMeta{
	LIVE: {AllowedArgs: []uint8{TDirIdCode}, IsArgTypeCode: false, TDirSize: DirSize, OpCode: 0x01, CyclesToExec: 10},
	LD: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x02,
		CyclesToExec: 5},
	ST:  {AllowedArgs: []uint8{TRegIdCode, TRegIdCode | TIndIdCode}, IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x03},
	ADD: {AllowedArgs: []uint8{TRegIdCode, TRegIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x04},
	SUB: {AllowedArgs: []uint8{TRegIdCode, TRegIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x05},
	AND: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x06},
	OR: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x07},
	XOR: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x08},
	ZJMP: {AllowedArgs: []uint8{TDirIdCode}, IsArgTypeCode: false, TDirSize: ShortDirSize, OpCode: 0x09},
	LDI: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: ShortDirSize, OpCode: 0x0A},
	STI: {AllowedArgs: []uint8{TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode},
		IsArgTypeCode: true, TDirSize: ShortDirSize, OpCode: 0x0B},
	FORK: {AllowedArgs: []uint8{TDirIdCode}, IsArgTypeCode: false, TDirSize: ShortDirSize, OpCode: 0x0C},
	LLD:  {AllowedArgs: []uint8{TDirIdCode | TIndIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x0D},
	LLDI: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: ShortDirSize, OpCode: 0x0E},
	LFORK: {AllowedArgs: []uint8{TDirIdCode}, IsArgTypeCode: false, TDirSize: ShortDirSize, OpCode: 0x0F},
	AFF:   {AllowedArgs: []uint8{TRegIdCode}, IsArgTypeCode: false, TDirSize: DirSize, OpCode: 0x10},
}
