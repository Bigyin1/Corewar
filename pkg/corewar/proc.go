package corewar

import (
	"corewar/pkg/consts"
)

func newProc(id, plID, initPos int, vm *VM) *proc {
	p := proc{id: id, vm: vm, pc: initPos}
	p.regs[0] = int32(-plID) // set 1st register to -playerID
	return &p
}

type proc struct {
	id        int
	carry     bool
	cmdMeta   consts.InstructionMeta
	liveCycle int
	execLeft  int
	pc        int
	b2nextOp  int
	regs      [consts.RegNumber]int32

	vm   *VM
	next *proc
}

func (p *proc) copy(pc int) {

	newProc := *p

	var newRegs [consts.RegNumber]int32
	copy(newRegs[:], p.regs[:])
	newProc.regs = newRegs
	newProc.pc = pc

	p.vm.procs.Put(&newProc)
}

func (p *proc) storeInReg(rIdx int, val int32) {
	p.regs[rIdx-1] = val
}

func (p *proc) loadFromReg(rIdx int) int32 {
	return p.regs[rIdx-1]
}

func (p *proc) loadArg(posArgs uint8, from arg) int32 {
	if posArgs&from.typ == 0 {
		panic("wrong val type")
	}
	if from.typ == consts.TDirIdCode {
		return int32(from.val)
	}

	if from.typ == consts.TRegIdCode {

		return p.loadFromReg(from.val)
	}
	if p.cmdMeta.IdxMod {
		from.val %= consts.IdxMod
	}

	return p.vm.field.GetInt32(p.pc + from.val)
}

func (p *proc) storeVal(posArgs uint8, to arg, val int32) {
	if posArgs&to.typ == 0 {
		panic("wrong val type")
	}
	if to.typ == consts.TRegIdCode {
		p.storeInReg(to.val, val)
	}

	if p.cmdMeta.IdxMod {
		to.val %= consts.IdxMod
	}
	p.vm.field.PutInt32(p.pc+to.val, val)
}

type arg struct {
	typ uint8
	val int
}

func (p *proc) live(args ...arg) {
	if len(args) != 1 {
		panic("live: wrong args count")
	}
	val1 := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	if val1 > 0 && val1 <= int32(len(p.vm.players)) {
		p.vm.lastAlive = &p.vm.players[val1-1]
	}

	p.liveCycle = p.vm.cyclesPassed + 1
	p.execLeft = p.cmdMeta.CyclesToExec
}

func (p *proc) ld(args ...arg) {
	defer func() {
		p.execLeft = p.cmdMeta.CyclesToExec
	}()
	if len(args) != 2 {
		panic("ld: wrong args count")
	}

	val1 := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	p.carry = false
	if val1 == 0 {
		p.carry = true
	}
	p.storeVal(p.cmdMeta.AllowedArgs[1], args[1], val1)
}

func (p *proc) st(args ...arg) {
	defer func() {
		p.execLeft = p.cmdMeta.CyclesToExec
	}()
	if len(args) != 2 {
		panic("st: wrong args count")
	}

	val1 := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	p.storeVal(p.cmdMeta.AllowedArgs[1], args[1], val1)

}

func (p *proc) add(args ...arg) {

	if len(args) != 3 {
		panic("add: wrong args count")
	}

	val1 := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArg(p.cmdMeta.AllowedArgs[1], args[1])
	p.storeVal(p.cmdMeta.AllowedArgs[2], args[3], val1+val2)
}

func (p *proc) sub(args ...arg) {

	if len(args) != 3 {
		panic("sub: wrong args count")
	}
	val1 := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArg(p.cmdMeta.AllowedArgs[1], args[1])
	p.storeVal(p.cmdMeta.AllowedArgs[2], args[3], val1-val2)
}

func (p *proc) and(args ...arg) {
	if len(args) != 3 {
		panic("and: wrong args count")
	}
	val1 := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArg(p.cmdMeta.AllowedArgs[1], args[1])
	res := val1 & val2
	p.carry = false
	if res == 0 {
		p.carry = true
	}
	p.storeVal(p.cmdMeta.AllowedArgs[2], args[2], res)
}

func (p *proc) or(args ...arg) {
	if len(args) != 3 {
		panic("and: wrong args count")
	}
	val1 := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArg(p.cmdMeta.AllowedArgs[1], args[1])
	res := val1 | val2
	p.carry = false
	if res == 0 {
		p.carry = true
	}
	p.storeVal(p.cmdMeta.AllowedArgs[2], args[2], res)
}

func (p *proc) xor(args ...arg) {
	if len(args) != 3 {
		panic("and: wrong args count")
	}
	val1 := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArg(p.cmdMeta.AllowedArgs[1], args[1])
	res := val1 ^ val2
	p.carry = false
	if res == 0 {
		p.carry = true
	}
	p.storeVal(p.cmdMeta.AllowedArgs[2], args[2], res)
}

func (p *proc) zjmp(args ...arg) {
	if len(args) != 1 {
		panic("zjmp: wrong args count")
	}
	val1 := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])

	if p.carry {
		p.pc += int(val1 % consts.IdxMod)
	}
}

func (p *proc) ldi(args ...arg) {
	if len(args) != 3 {
		panic("ldi: wrong args count")
	}

	val1 := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArg(p.cmdMeta.AllowedArgs[1], args[1])
	addr := int32(p.pc) + (val1+val2)%consts.IdxMod
	p.storeVal(p.cmdMeta.AllowedArgs[2], args[2], p.vm.field.GetInt32(int(addr)))
}

func (p *proc) sti(args ...arg) {
	if len(args) != 3 {
		panic("sti: wrong args count")
	}

	storeVal := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	val1 := p.loadArg(p.cmdMeta.AllowedArgs[1], args[1])
	val2 := p.loadArg(p.cmdMeta.AllowedArgs[2], args[2])

	addr := int32(p.pc) + (val1+val2)%consts.IdxMod
	p.vm.field.PutInt32(int(addr), storeVal)
}

func (p *proc) fork(args ...arg) {
	pc := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	p.copy(int(pc) % consts.IdxMod)
}

func (p *proc) lld(args ...arg) {
	defer func() {
		p.execLeft = p.cmdMeta.CyclesToExec
	}()
	if len(args) != 2 {
		panic("lld: wrong args count")
	}

	val1 := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	p.carry = false
	if val1 == 0 {
		p.carry = true
	}
	p.storeVal(p.cmdMeta.AllowedArgs[1], args[1], val1)
}

func (p *proc) lldi(args ...arg) {
	if len(args) != 3 {
		panic("ldi: wrong args count")
	}

	val1 := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArg(p.cmdMeta.AllowedArgs[1], args[1])
	addr := int32(p.pc) + (val1 + val2)
	p.storeVal(p.cmdMeta.AllowedArgs[2], args[2], p.vm.field.GetInt32(int(addr)))
}

func (p *proc) lfork(args ...arg) {
	pc := p.loadArg(p.cmdMeta.AllowedArgs[0], args[0])
	p.copy(int(pc))
}
