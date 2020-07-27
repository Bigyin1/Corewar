package consts

type InstructionMeta struct {
	AllowedArgs   []uint8
	IsArgTypeCode bool
	TDirSize      int
	OpCode        byte
	CyclesToExec  int
}

var InstructionsConfig = map[InstructionName]InstructionMeta{
	LIVE: {AllowedArgs: []uint8{TDirIdCode}, IsArgTypeCode: false, TDirSize: DirSize, OpCode: 0x01, CyclesToExec: 10},
	LD: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x02,
		CyclesToExec: 5},
	ST: {AllowedArgs: []uint8{TRegIdCode, TRegIdCode | TIndIdCode}, IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x03,
		CyclesToExec: 5},
	ADD: {AllowedArgs: []uint8{TRegIdCode, TRegIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x04, CyclesToExec: 10},
	SUB: {AllowedArgs: []uint8{TRegIdCode, TRegIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x05, CyclesToExec: 10},
	AND: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x06, CyclesToExec: 6},
	OR: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x07, CyclesToExec: 6},
	XOR: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x08, CyclesToExec: 6},
	ZJMP: {AllowedArgs: []uint8{TDirIdCode}, IsArgTypeCode: false, TDirSize: ShortDirSize, OpCode: 0x09, CyclesToExec: 20},
	LDI: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: ShortDirSize, OpCode: 0x0A, CyclesToExec: 25},
	STI: {AllowedArgs: []uint8{TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode},
		IsArgTypeCode: true, TDirSize: ShortDirSize, OpCode: 0x0B, CyclesToExec: 25},
	FORK: {AllowedArgs: []uint8{TDirIdCode}, IsArgTypeCode: false, TDirSize: ShortDirSize, OpCode: 0x0C, CyclesToExec: 800},
	LLD:  {AllowedArgs: []uint8{TDirIdCode | TIndIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirSize, OpCode: 0x0D, CyclesToExec: 10},
	LLDI: {AllowedArgs: []uint8{TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: ShortDirSize, OpCode: 0x0E, CyclesToExec: 50},
	LFORK: {AllowedArgs: []uint8{TDirIdCode}, IsArgTypeCode: false, TDirSize: ShortDirSize, OpCode: 0x0F, CyclesToExec: 1000},
	AFF:   {AllowedArgs: []uint8{TRegIdCode}, IsArgTypeCode: false, TDirSize: DirSize, OpCode: 0x10},
}
