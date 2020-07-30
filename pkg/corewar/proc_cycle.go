package corewar

import (
	"corewar/pkg/consts"
	"fmt"
)

type instr struct {
	meta consts.InstructionMeta
	f    func(p *proc, a ...arg)
}

// TODO refactor logic to remove this table
var opcodeToInstr = []instr{
	{meta: consts.InstructionsConfig[consts.LIVE], f: Live},
	{meta: consts.InstructionsConfig[consts.LD], f: Ld},
	{meta: consts.InstructionsConfig[consts.ST], f: St},
	{meta: consts.InstructionsConfig[consts.ADD], f: Add},
	{meta: consts.InstructionsConfig[consts.SUB], f: Sub},
	{meta: consts.InstructionsConfig[consts.AND], f: And},
	{meta: consts.InstructionsConfig[consts.OR], f: Or},
	{meta: consts.InstructionsConfig[consts.XOR], f: Xor},
	{meta: consts.InstructionsConfig[consts.ZJMP], f: Zjmp},
	{meta: consts.InstructionsConfig[consts.LDI], f: Ldi},
	{meta: consts.InstructionsConfig[consts.STI], f: Sti},
	{meta: consts.InstructionsConfig[consts.FORK], f: Fork},
	{meta: consts.InstructionsConfig[consts.LLD], f: Lld},
	{meta: consts.InstructionsConfig[consts.LLDI], f: Lldi},
	{meta: consts.InstructionsConfig[consts.LFORK], f: Lfork},
	{meta: consts.InstructionsConfig[consts.AFF], f: Aff},
}

func (p *proc) getArgSize(tid consts.TypeID) int {
	switch tid {
	case consts.TDirIdCode:
		return p.opMeta.TDirSize
	case consts.TRegIdCode:
		return consts.RegArgSize
	case consts.TIndIdCode:
		return consts.IndArgSize
	}
	return 0
}

func (p *proc) shiftToNextOp(argTypes []consts.TypeID) {
	p.pc += 1 //opcode
	if p.opMeta.IsArgTypeCode {
		p.pc += 1
	}
	for _, at := range argTypes {
		p.pc += p.getArgSize(at)
	}
}

func (p *proc) shiftToNextOp2(args []arg) {
	p.pc += 1 //opcode
	if p.opMeta.IsArgTypeCode {
		p.pc += 1
	}
	for _, a := range args {
		p.pc += p.getArgSize(a.typ)
	}
}

func (p *proc) parseArgsTypes() ([]consts.TypeID, bool) {
	if !p.opMeta.IsArgTypeCode {
		return p.opMeta.AllowedArgs, true
	}

	var ok = true

	var expArgs []consts.TypeID
	argTypeCode := p.vm.field.getByte(p.pc + 1)
	offset := 6
	toLeft := 0
	for _, aa := range p.opMeta.AllowedArgs {
		var byteCode byte
		byteCode |= argTypeCode
		byteCode <<= toLeft
		byteCode >>= offset
		if byteCode == 0 {
			ok = false
			return nil, ok
		}
		argType := consts.ByteCodeToTypeID(byteCode)
		if argType&aa == 0 {
			ok = false
		}
		expArgs = append(expArgs, argType)
		toLeft += 2
	}
	return expArgs, ok
}

func (p *proc) parseArgValues(argTypes []consts.TypeID) (args []arg, ok bool) {
	var offset = 1
	if p.opMeta.IsArgTypeCode {
		offset++
	}
	ok = true
	for _, at := range argTypes {
		switch at {
		case consts.TRegIdCode:
			rv := p.vm.field.getByte(p.pc + offset)
			offset += consts.RegArgSize
			if rv <= 0 || rv > consts.RegNumber {
				ok = false
				return
			}
			args = append(args, arg{val: int(rv), typ: at})
		case consts.TDirIdCode:
			var dv = p.vm.field.getInt32(p.pc + offset)
			if p.opMeta.TDirSize == consts.ShortDirSize {
				dv = p.vm.field.getInt16(p.pc + offset)
			}
			offset += p.opMeta.TDirSize
			args = append(args, arg{val: dv, typ: at})
		case consts.TIndIdCode:
			args = append(args, arg{val: p.vm.field.getInt16(p.pc + offset), typ: at})
			offset += consts.IndArgSize
		}
	}
	return
}

func (p *proc) getOpArgs() ([]arg, bool) {
	var args []arg
	argTypes, ok := p.parseArgsTypes()
	if !ok {
		p.shiftToNextOp(argTypes)
		return nil, false
	}
	args, ok = p.parseArgValues(argTypes)
	if !ok {
		p.shiftToNextOp(argTypes)
		return nil, false
	}
	return args, true
}

func (p *proc) setOpCode() {
	if p.execLeft != 0 {
		return
	}
	p.currOpCode = p.vm.field.getByte(p.pc)
	if p.currOpCode <= 0x10 && p.currOpCode >= 0x01 {
		op := opcodeToInstr[p.currOpCode-1]
		p.execLeft = op.meta.CyclesToExec
	}

}

func (p *proc) execOp() {
	if p.execLeft != 0 {
		return
	}
	if p.currOpCode > 0x10 || p.currOpCode < 0x01 {
		p.pc++
		return
	}
	op := opcodeToInstr[p.currOpCode-1]
	p.opMeta = op.meta

	args, ok := p.getOpArgs()
	if !ok {
		return
	}
	op.f(p, args...)
}

func (p *proc) Cycle() {
	if p.vm.currCycle == 109 && p.id == 0 {
		fmt.Println()
	}
	if p.id == 0 && p.vm.currCycle == 104 {
		fmt.Print()
	}
	p.setOpCode()
	if p.execLeft > 0 {
		p.execLeft--
	}
	p.execOp()
}
