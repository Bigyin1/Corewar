package consts

const NameHeader = ".name"
const ChampNameLength = 64
const CommentHeader = ".comment"
const LabelChars = "abcdefghijklmnopqrstuvwxyz_0123456789"
const CommentSymbol = "#"
const RegNumber = 16
const RegSize = 4
const SeparatorSymbol = ","

const MagicHeader = "\x00\xea\x83\xf3"
const NullSeq = "\x00\x00\x00\x00"

const ProgNameLength = 128
const CommentLength = 2048

const RegArgSize = 1
const ShortDirSize = 2
const IndArgSize = 2
const DirArgSize = 4

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

type TypeID uint8

// Type codes for internal usage, not for byte code
const (
	TRegIdCode TypeID = 1 << 0
	TDirIdCode TypeID = 1 << 1
	TIndIdCode TypeID = 1 << 2
)

var typeArr = []TypeID{TRegIdCode, TDirIdCode, TIndIdCode}

func GetArgSize(m InstructionMeta, tid TypeID) int {
	switch tid {
	case TDirIdCode:
		return m.TDirSize
	case TRegIdCode:
		return RegArgSize
	case TIndIdCode:
		return IndArgSize
	}
	return 0
}

func ByteCodeToTypeID(bc byte) TypeID {
	if bc-1 < 0 || int(bc-1) >= len(typeArr) {
		panic("error")
	}
	return typeArr[bc-1]
}

func TypeIDToByteCode(tid TypeID) byte {
	switch tid {
	case TRegIdCode:
		return 0b01
	case TDirIdCode:
		return 0b10
	case TIndIdCode:
		return 0b11
	}
	return 0
}
