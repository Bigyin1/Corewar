package corewar

import "corewar/pkg/consts"

type instr struct {
	meta consts.InstructionMeta
	f    func(p *proc, a ...arg)
}

var opcodeToInstr = map[byte]instr{
	0x01: {meta: consts.InstructionsConfig[consts.LIVE], f: Live},
	0x02: {meta: consts.InstructionsConfig[consts.LD], f: Ld},
	0x03: {meta: consts.InstructionsConfig[consts.ST], f: St},
	0x04: {meta: consts.InstructionsConfig[consts.ADD], f: Add},
	0x05: {meta: consts.InstructionsConfig[consts.SUB], f: Sub},
	0x06: {meta: consts.InstructionsConfig[consts.AND], f: And},
	0x07: {meta: consts.InstructionsConfig[consts.OR], f: Or},
	0x08: {meta: consts.InstructionsConfig[consts.XOR], f: Xor},
	0x09: {meta: consts.InstructionsConfig[consts.ZJMP], f: Zjmp},
	0x0A: {meta: consts.InstructionsConfig[consts.LDI], f: Ldi},
	0x0B: {meta: consts.InstructionsConfig[consts.STI], f: Sti},
	0x0C: {meta: consts.InstructionsConfig[consts.FORK], f: Fork},
	0x0D: {meta: consts.InstructionsConfig[consts.LLD], f: Lld},
	0x0E: {meta: consts.InstructionsConfig[consts.LLDI], f: Lldi},
	0x0F: {meta: consts.InstructionsConfig[consts.LFORK], f: Lfork},
	0x10: {meta: consts.InstructionsConfig[consts.AFF], f: Aff},
}

func (p *proc) setOpCode() {
	if p.execLeft != 0 {
		return
	}
	var opCode [1]byte
	p.vm.field.loadFrom(p.pc, opCode[:])
	p.currOpCode = opCode[0]
}

func (p *proc) execOp() {

}

func (p *proc) Cycle() {
	p.setOpCode()
	if p.execLeft > 0 {
		p.execLeft--
	}

}
