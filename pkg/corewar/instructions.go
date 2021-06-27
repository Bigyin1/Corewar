package corewar

import (
	"fmt"

	"github.com/Bigyin1/Corewar/pkg/consts"
)

type arg struct {
	typ consts.TypeID
	val int
}

func Live(p *proc, args ...arg) {
	if len(args) != 1 {
		panic("Live: wrong args count")
	}
	pID := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	if -pID > 0 && -pID <= len(p.vm.players) {
		p.vm.lastAlive = &p.vm.players[-pID-1]
	}

	p.vm.liveOps++
	p.liveCycle = p.vm.currCycle

	if p.vm.log {
		fmt.Printf("P\t%d | %s %d\n", p.id, "live", pID)
	}
	p.shiftToNextOp2(args)
}

func Ld(p *proc, args ...arg) {
	if len(args) != 2 {
		panic("Ld: wrong args count")
	}

	val := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	p.carry = false
	if val == 0 {
		p.carry = true
	}
	p.storeValToArg(p.opMeta.AllowedArgs[1], args[1], val)
	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "ld")
		if args[0].typ == consts.TDirIdCode {
			fmt.Printf("%d r%d\n", val, args[1].val)
		} else {
			fmt.Printf("from addr: %d val: %d r%d\n", args[0].val, val, args[1].val)
		}
	}
	p.shiftToNextOp2(args)
}

func St(p *proc, args ...arg) {
	if len(args) != 2 {
		panic("St: wrong args count")
	}

	regVal := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	p.storeValToArg(p.opMeta.AllowedArgs[1], args[1], regVal)

	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "st")
		if args[1].typ == consts.TRegIdCode {
			fmt.Printf("r%d(%d) r%d\n", args[0].val, regVal, args[1].val)
		} else {
			fmt.Printf("r%d(%d) %d\n", args[0].val, regVal, args[1].val)
		}
	}
	p.shiftToNextOp2(args)

}

func Add(p *proc, args ...arg) {

	if len(args) != 3 {
		panic("Add: wrong args count")
	}

	val1 := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.opMeta.AllowedArgs[1], args[1])
	p.storeValToArg(p.opMeta.AllowedArgs[2], args[2], val1+val2)
	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "add")
		fmt.Printf("r%d(%d) r%d(%d) r%d\n",
			args[0].val, val1,
			args[1].val, val2,
			args[2].val)
	}
	p.shiftToNextOp2(args)
}

func Sub(p *proc, args ...arg) {

	if len(args) != 3 {
		panic("Sub: wrong args count")
	}
	val1 := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.opMeta.AllowedArgs[1], args[1])
	p.storeValToArg(p.opMeta.AllowedArgs[2], args[2], val1-val2)
	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "sub")
		fmt.Printf("r%d(%d) r%d(%d) r%d\n",
			args[0].val, val1,
			args[1].val, val2,
			args[2].val)
	}
	p.shiftToNextOp2(args)
}

func And(p *proc, args ...arg) {
	if len(args) != 3 {
		panic("And: wrong args count")
	}
	val1 := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.opMeta.AllowedArgs[1], args[1])
	res := val1 & val2
	p.carry = false
	if res == 0 {
		p.carry = true
	}
	p.storeValToArg(p.opMeta.AllowedArgs[2], args[2], res)
	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "and")
		fmt.Printf("r%d(%d) r%d(%d) r%d\n",
			args[0].val, val1,
			args[1].val, val2,
			args[2].val)
	}
	p.shiftToNextOp2(args)
}

func Or(p *proc, args ...arg) {
	if len(args) != 3 {
		panic("And: wrong args count")
	}
	val1 := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.opMeta.AllowedArgs[1], args[1])
	res := val1 | val2
	p.carry = false
	if res == 0 {
		p.carry = true
	}
	p.storeValToArg(p.opMeta.AllowedArgs[2], args[2], res)
	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "or")
		fmt.Printf("r%d(%d) r%d(%d) r%d\n",
			args[0].val, val1,
			args[1].val, val2,
			args[2].val)
	}
	p.shiftToNextOp2(args)
}

func Xor(p *proc, args ...arg) {
	if len(args) != 3 {
		panic("And: wrong args count")
	}
	val1 := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.opMeta.AllowedArgs[1], args[1])
	res := val1 ^ val2
	p.carry = false
	if res == 0 {
		p.carry = true
	}
	p.storeValToArg(p.opMeta.AllowedArgs[2], args[2], res)
	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "xor")
		fmt.Printf("r%d(%d) r%d(%d) r%d\n",
			args[0].val, val1,
			args[1].val, val2,
			args[2].val)
	}
	p.shiftToNextOp2(args)
}

func Zjmp(p *proc, args ...arg) {
	if len(args) != 1 {
		panic("Zjmp: wrong args count")
	}
	val1 := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])

	if p.carry {
		p.pc += val1 % consts.IdxMod
	} else {
		p.shiftToNextOp2(args)
	}

	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "zjmp")
		fmt.Printf("%d %t\n", val1, p.carry)
	}
}

func Ldi(p *proc, args ...arg) {
	if len(args) != 3 {
		panic("Ldi: wrong args count")
	}

	val1 := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.opMeta.AllowedArgs[1], args[1])
	addr := p.pc + (val1+val2)%consts.IdxMod
	p.storeValToArg(p.opMeta.AllowedArgs[2], args[2], p.vm.field.getInt32(addr))

	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "ldi")
		fmt.Printf("%d %d r%d\n", val1, val2, args[2].val)
		fmt.Printf("\t\t| -> loaded from %d\n", addr)
	}
	p.shiftToNextOp2(args)
}

func Sti(p *proc, args ...arg) {
	if len(args) != 3 {
		panic("Sti: wrong args count")
	}

	storeVal := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	val1 := p.loadArgVal(p.opMeta.AllowedArgs[1], args[1])
	val2 := p.loadArgVal(p.opMeta.AllowedArgs[2], args[2])

	addr := p.pc + (val1+val2)%consts.IdxMod
	p.vm.field.putInt32(addr, storeVal)
	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "sti")
		fmt.Printf("r%d(%d) %d %d\n", args[0].val, storeVal, val1, val2)
		fmt.Printf("\t\t| -> stored to %d\n", addr)
	}
	p.shiftToNextOp2(args)
}

func Fork(p *proc, args ...arg) {
	addr := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	p.copy(p.pc + addr%consts.IdxMod)
	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "fork")
		fmt.Printf("%d (%d)\n", addr, p.pc+addr%consts.IdxMod)
	}
	p.shiftToNextOp2(args)
}

func Lld(p *proc, args ...arg) {
	if len(args) != 2 {
		panic("lld: wrong args count")
	}

	val := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0], true)
	p.carry = false
	if val == 0 {
		p.carry = true
	}
	p.storeValToArg(p.opMeta.AllowedArgs[1], args[1], val)
	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "lld")
		if args[0].typ == consts.TDirIdCode {
			fmt.Printf("%d r%d\n", val, args[1].val)
		} else {
			fmt.Printf("from addr: %d val: %d r%d\n", args[0].val, val, args[1].val)
		}
	}
	p.shiftToNextOp2(args)
}

func Lldi(p *proc, args ...arg) {
	if len(args) != 3 {
		panic("Ldi: wrong args count")
	}

	val1 := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	val2 := p.loadArgVal(p.opMeta.AllowedArgs[1], args[1])
	addr := p.pc + (val1 + val2)
	p.storeValToArg(p.opMeta.AllowedArgs[2], args[2], p.vm.field.getInt32(addr))
	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "lldi")
		fmt.Printf("%d %d r%d ", val1, val2, args[2].val)
		fmt.Printf("loaded from %d\n", addr)
	}
	p.shiftToNextOp2(args)
}

func Lfork(p *proc, args ...arg) {
	addr := p.loadArgVal(p.opMeta.AllowedArgs[0], args[0])
	p.copy(p.pc + addr)
	if p.vm.log {
		fmt.Printf("P\t%d | %s ", p.id, "lfork")
		fmt.Printf("%d (%d)\n", addr, p.pc+addr)
	}
	p.shiftToNextOp2(args)
}

func Aff(p *proc, args ...arg) {
	p.shiftToNextOp2(args)
}
