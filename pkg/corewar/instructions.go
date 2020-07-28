package corewar

import "corewar/pkg/consts"

type arg struct {
	typ uint8
	val int
}

func Live(p *proc, args ...arg) {
	if len(args) != 1 {
		panic("Live: wrong args count")
	}
	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	if val1 > 0 && val1 <= len(p.vm.players) {
		p.vm.lastAlive = &p.vm.players[val1-1]
	}

	p.liveCycle = p.vm.cyclesPassed + 1
	p.execLeft = p.cmdMeta.CyclesToExec
}

func Ld(p *proc, args ...arg) {
	defer func() {
		p.execLeft = p.cmdMeta.CyclesToExec
	}()
	if len(args) != 2 {
		panic("Ld: wrong args count")
	}

	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	p.carry = false
	if val1 == 0 {
		p.carry = true
	}
	p.storeValToArg(p.cmdMeta.AllowedArgs[1], args[1], val1)
}

func St(p *proc, args ...arg) {
	defer func() {
		p.execLeft = p.cmdMeta.CyclesToExec
	}()
	if len(args) != 2 {
		panic("St: wrong args count")
	}

	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	p.storeValToArg(p.cmdMeta.AllowedArgs[1], args[1], val1)

}

func Add(p *proc, args ...arg) {

	if len(args) != 3 {
		panic("Add: wrong args count")
	}

	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	p.storeValToArg(p.cmdMeta.AllowedArgs[2], args[2], val1+val2)
}

func Sub(p *proc, args ...arg) {

	if len(args) != 3 {
		panic("Sub: wrong args count")
	}
	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	p.storeValToArg(p.cmdMeta.AllowedArgs[2], args[2], val1-val2)
}

func And(p *proc, args ...arg) {
	if len(args) != 3 {
		panic("And: wrong args count")
	}
	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	res := val1 & val2
	p.carry = false
	if res == 0 {
		p.carry = true
	}
	p.storeValToArg(p.cmdMeta.AllowedArgs[2], args[2], res)
}

func Or(p *proc, args ...arg) {
	if len(args) != 3 {
		panic("And: wrong args count")
	}
	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	res := val1 | val2
	p.carry = false
	if res == 0 {
		p.carry = true
	}
	p.storeValToArg(p.cmdMeta.AllowedArgs[2], args[2], res)
}

func Xor(p *proc, args ...arg) {
	if len(args) != 3 {
		panic("And: wrong args count")
	}
	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	res := val1 ^ val2
	p.carry = false
	if res == 0 {
		p.carry = true
	}
	p.storeValToArg(p.cmdMeta.AllowedArgs[2], args[2], res)
}

func Zjmp(p *proc, args ...arg) {
	if len(args) != 1 {
		panic("Zjmp: wrong args count")
	}
	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])

	if p.carry {
		p.pc += val1 % consts.IdxMod
	}
}

func Ldi(p *proc, args ...arg) {
	if len(args) != 3 {
		panic("Ldi: wrong args count")
	}

	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	addr := p.pc + (val1+val2)%consts.IdxMod
	p.storeValToArg(p.cmdMeta.AllowedArgs[2], args[2], p.vm.field.getInt32(addr))
}

func Sti(p *proc, args ...arg) {
	if len(args) != 3 {
		panic("Sti: wrong args count")
	}

	storeVal := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[2], args[2])

	addr := p.pc + (val1+val2)%consts.IdxMod
	p.vm.field.putInt32(addr, storeVal)
}

func Fork(p *proc, args ...arg) {
	pc := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	p.copy(pc % consts.IdxMod)
}

func Lld(p *proc, args ...arg) {
	defer func() {
		p.execLeft = p.cmdMeta.CyclesToExec
	}()
	if len(args) != 2 {
		panic("lld: wrong args count")
	}

	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0], true)
	p.carry = false
	if val1 == 0 {
		p.carry = true
	}
	p.storeValToArg(p.cmdMeta.AllowedArgs[1], args[1], val1)
}

func Lldi(p *proc, args ...arg) {
	if len(args) != 3 {
		panic("Ldi: wrong args count")
	}

	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	addr := p.pc + (val1 + val2)
	p.storeValToArg(p.cmdMeta.AllowedArgs[2], args[2], p.vm.field.getInt32(addr))
}

func Lfork(p *proc, args ...arg) {
	pc := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	p.copy(pc)
}

func Aff(p *proc, args ...arg) {}
