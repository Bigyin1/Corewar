package corewar

import "corewar/pkg/consts"

type arg struct {
	typ uint8
	val int
}

func (p *proc) Live(args ...arg) {
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

func (p *proc) Ld(args ...arg) {
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

func (p *proc) St(args ...arg) {
	defer func() {
		p.execLeft = p.cmdMeta.CyclesToExec
	}()
	if len(args) != 2 {
		panic("St: wrong args count")
	}

	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	p.storeValToArg(p.cmdMeta.AllowedArgs[1], args[1], val1)

}

func (p *proc) Add(args ...arg) {

	if len(args) != 3 {
		panic("Add: wrong args count")
	}

	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	p.storeValToArg(p.cmdMeta.AllowedArgs[2], args[2], val1+val2)
}

func (p *proc) Sub(args ...arg) {

	if len(args) != 3 {
		panic("Sub: wrong args count")
	}
	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	p.storeValToArg(p.cmdMeta.AllowedArgs[2], args[2], val1-val2)
}

func (p *proc) And(args ...arg) {
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

func (p *proc) Or(args ...arg) {
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

func (p *proc) Xor(args ...arg) {
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

func (p *proc) Zjmp(args ...arg) {
	if len(args) != 1 {
		panic("Zjmp: wrong args count")
	}
	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])

	if p.carry {
		p.pc += val1 % consts.IdxMod
	}
}

func (p *proc) Ldi(args ...arg) {
	if len(args) != 3 {
		panic("Ldi: wrong args count")
	}

	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	addr := p.pc + (val1+val2)%consts.IdxMod
	p.storeValToArg(p.cmdMeta.AllowedArgs[2], args[2], p.vm.field.GetInt32(addr))
}

func (p *proc) Sti(args ...arg) {
	if len(args) != 3 {
		panic("Sti: wrong args count")
	}

	storeVal := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[2], args[2])

	addr := p.pc + (val1+val2)%consts.IdxMod
	p.vm.field.PutInt32(addr, storeVal)
}

func (p *proc) Fork(args ...arg) {
	pc := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	p.copy(pc % consts.IdxMod)
}

func (p *proc) lld(args ...arg) {
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

func (p *proc) Lldi(args ...arg) {
	if len(args) != 3 {
		panic("Ldi: wrong args count")
	}

	val1 := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.cmdMeta.AllowedArgs[1], args[1])
	addr := p.pc + (val1 + val2)
	p.storeValToArg(p.cmdMeta.AllowedArgs[2], args[2], p.vm.field.GetInt32(addr))
}

func (p *proc) Lfork(args ...arg) {
	pc := p.loadArgVal(p.cmdMeta.AllowedArgs[0], args[0])
	p.copy(pc)
}
