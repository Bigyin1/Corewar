package consts

type InstructionMeta struct {
	AllowedArgs   []TypeID
	IsArgTypeCode bool
	TDirSize      int
	OpCode        byte
	CyclesToExec  int
}

var InstructionsConfig = map[InstructionName]InstructionMeta{
	LIVE: {AllowedArgs: []TypeID{TDirIdCode}, IsArgTypeCode: false, TDirSize: DirArgSize, OpCode: 0x01, CyclesToExec: 10},
	LD: {AllowedArgs: []TypeID{TDirIdCode | TIndIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x02,
		CyclesToExec: 5},
	ST: {AllowedArgs: []TypeID{TRegIdCode, TRegIdCode | TIndIdCode}, IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x03,
		CyclesToExec: 5},
	ADD: {AllowedArgs: []TypeID{TRegIdCode, TRegIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x04, CyclesToExec: 10},
	SUB: {AllowedArgs: []TypeID{TRegIdCode, TRegIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x05, CyclesToExec: 10},
	AND: {AllowedArgs: []TypeID{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x06, CyclesToExec: 6},
	OR: {AllowedArgs: []TypeID{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x07, CyclesToExec: 6},
	XOR: {AllowedArgs: []TypeID{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x08, CyclesToExec: 6},
	ZJMP: {AllowedArgs: []TypeID{TDirIdCode}, IsArgTypeCode: false, TDirSize: ShortDirSize, OpCode: 0x09, CyclesToExec: 20},
	LDI: {AllowedArgs: []TypeID{TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: ShortDirSize, OpCode: 0x0A, CyclesToExec: 25},
	STI: {AllowedArgs: []TypeID{TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode},
		IsArgTypeCode: true, TDirSize: ShortDirSize, OpCode: 0x0B, CyclesToExec: 25},
	FORK: {AllowedArgs: []TypeID{TDirIdCode}, IsArgTypeCode: false, TDirSize: ShortDirSize, OpCode: 0x0C, CyclesToExec: 800},
	LLD:  {AllowedArgs: []TypeID{TDirIdCode | TIndIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x0D, CyclesToExec: 10},
	LLDI: {AllowedArgs: []TypeID{TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: ShortDirSize, OpCode: 0x0E, CyclesToExec: 50},
	LFORK: {AllowedArgs: []TypeID{TDirIdCode}, IsArgTypeCode: false, TDirSize: ShortDirSize, OpCode: 0x0F, CyclesToExec: 1000},
	AFF:   {AllowedArgs: []TypeID{TRegIdCode}, IsArgTypeCode: false, TDirSize: DirArgSize, OpCode: 0x10},
}
