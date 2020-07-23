package corewar

import "corewar/pkg/consts"

type proc struct {
	id        int
	carry     bool
	opcode    byte
	liveCycle int
	execLeft  int
	currPos   int
	b2nextOp  int
	regs      [consts.RegNumber]int32

	f    *field
	next *proc
}

func newProc(id, plID, currPos int, f *field) *proc {
	p := proc{id: id, f: f, currPos: currPos}
	p.regs[0] = int32(-plID) // set 1st register to -playerID
	return &p
}
