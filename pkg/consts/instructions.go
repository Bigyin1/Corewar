package consts

type InstructionMeta struct {
	Name          InstructionName
	AllowedArgs   []TypeID
	IsArgTypeCode bool
	TDirSize      int
	OpCode        byte
	CyclesToExec  int
}
type InstrConfig []InstructionMeta

var InstructionsConfig = InstrConfig{
	{Name: LIVE, AllowedArgs: []TypeID{TDirIdCode}, IsArgTypeCode: false, TDirSize: DirArgSize, OpCode: 0x01, CyclesToExec: 10},
	{Name: LD, AllowedArgs: []TypeID{TDirIdCode | TIndIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x02,
		CyclesToExec: 5},
	{Name: ST, AllowedArgs: []TypeID{TRegIdCode, TRegIdCode | TIndIdCode}, IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x03,
		CyclesToExec: 5},
	{Name: ADD, AllowedArgs: []TypeID{TRegIdCode, TRegIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x04, CyclesToExec: 10},
	{Name: SUB, AllowedArgs: []TypeID{TRegIdCode, TRegIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x05, CyclesToExec: 10},
	{Name: AND, AllowedArgs: []TypeID{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x06, CyclesToExec: 6},
	{Name: OR, AllowedArgs: []TypeID{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x07, CyclesToExec: 6},
	{Name: XOR, AllowedArgs: []TypeID{TDirIdCode | TIndIdCode | TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x08, CyclesToExec: 6},
	{Name: ZJMP, AllowedArgs: []TypeID{TDirIdCode}, IsArgTypeCode: false, TDirSize: ShortDirSize, OpCode: 0x09, CyclesToExec: 20},
	{Name: LDI, AllowedArgs: []TypeID{TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: ShortDirSize, OpCode: 0x0A, CyclesToExec: 25},
	{Name: STI, AllowedArgs: []TypeID{TRegIdCode, TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode},
		IsArgTypeCode: true, TDirSize: ShortDirSize, OpCode: 0x0B, CyclesToExec: 25},
	{Name: FORK, AllowedArgs: []TypeID{TDirIdCode}, IsArgTypeCode: false, TDirSize: ShortDirSize, OpCode: 0x0C, CyclesToExec: 800},
	{Name: LLD, AllowedArgs: []TypeID{TDirIdCode | TIndIdCode, TRegIdCode}, IsArgTypeCode: true, TDirSize: DirArgSize, OpCode: 0x0D, CyclesToExec: 10},
	{Name: LLDI, AllowedArgs: []TypeID{TDirIdCode | TIndIdCode | TRegIdCode, TRegIdCode | TDirIdCode, TRegIdCode},
		IsArgTypeCode: true, TDirSize: ShortDirSize, OpCode: 0x0E, CyclesToExec: 50},
	{Name: LFORK, AllowedArgs: []TypeID{TDirIdCode}, IsArgTypeCode: false, TDirSize: ShortDirSize, OpCode: 0x0F, CyclesToExec: 1000},
	{Name: AFF, AllowedArgs: []TypeID{TRegIdCode}, IsArgTypeCode: false, TDirSize: DirArgSize, OpCode: 0x10},
}

func (ic *InstrConfig) FindByName(iName InstructionName) InstructionMeta {
	for i := range InstructionsConfig {
		cfg := InstructionsConfig[i]
		if iName == cfg.Name {
			return cfg
		}
	}
	return InstructionMeta{}
}
