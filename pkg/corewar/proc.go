package corewar

import (
	"corewar/pkg/consts"
)

func newProc(id, plID, initPos int, vm *VM) *proc {
	p := proc{id: id, vm: vm, pc: initPos}
	p.regs[0] = -plID // set 1st register to -playerID
	return &p
}

type proc struct {
	id        int
	liveCycle int

	regs       [consts.RegNumber]int
	execLeft   int
	pc         int
	carry      bool
	currOpCode byte
	opMeta     consts.InstructionMeta

	vm   *VM
	next *proc
}

func (p *proc) copy(pc int) {

	newProc := *p

	var newRegs [consts.RegNumber]int
	copy(newRegs[:], p.regs[:])
	newProc.regs = newRegs
	newProc.pc = pc

	p.vm.procs.Put(&newProc)
}

func (p *proc) storeReg(rIdx int, val int) {
	p.regs[rIdx-1] = val
}

func (p *proc) loadReg(rIdx int) int {
	return p.regs[rIdx-1]
}

func (p *proc) loadArgVal(posArgs consts.TypeID, from arg, nomod ...bool) int {
	if posArgs&from.typ == 0 {
		panic("wrong val type")
	}
	if from.typ == consts.TDirIdCode {
		return from.val
	}

	if from.typ == consts.TRegIdCode {
		return p.loadReg(from.val)
	}
	if len(nomod) == 0 {
		from.val %= consts.IdxMod
	}

	return p.vm.field.getInt32(p.pc + from.val)
}

func (p *proc) storeValToArg(posArgs consts.TypeID, to arg, val int) {
	if posArgs&to.typ == 0 {
		panic("wrong val type")
	}
	if to.typ == consts.TRegIdCode {
		p.storeReg(to.val, val)
	}

	to.val %= consts.IdxMod
	p.vm.field.putInt32(p.pc+to.val, val)
}
