package corewar

import (
	"corewar/pkg/consts"
	"reflect"
	"testing"
)

func TestInstructions(t *testing.T) {

	tvm := newTestVM()

	t.Run("Live", tvm.testLive)
	t.Run("Ld", tvm.testLD)
	t.Run("St", tvm.testST)
	t.Run("Add", tvm.testAdd)
	t.Run("Sub", tvm.testSub)
	t.Run("And", tvm.testAnd)
	t.Run("Or", tvm.testOr)
	t.Run("Xor", tvm.testXor)
	t.Run("Zjmp", tvm.testZjmp)
	t.Run("Ldi", tvm.testLdi)
	t.Run("Sti", tvm.testSti)
	t.Run("Fork", tvm.testFork)
	t.Run("Lld", tvm.testLLD)
	t.Run("Llld", tvm.testLldi)
	t.Run("Lfork", tvm.testLfork)

}

type testVM struct {
	vm *VM
	p  *proc
}

func newTestVM() *testVM {
	playerID := 1
	pc := 0
	fieldSz := 33

	vm := NewVM()
	vm.field = newField(fieldSz)
	vm.players = []player{{id: 1}, {id: 2}, {id: 3}}
	testProc := newProc(1, playerID, pc, vm)
	vm.procs.Put(testProc)

	return &testVM{
		vm: vm,
		p:  testProc,
	}

}

func (tvm *testVM) testLive(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.LIVE]
	arg := arg{consts.TDirIdCode, 1}
	Live(tvm.p, arg)
	if tvm.vm.lastAlive.id != arg.val {
		t.Errorf("wrong last alive player")
		return
	}
	if tvm.p.liveCycle != 1 {
		t.Errorf("wrong alive player")
		return
	}

	*tvm = *newTestVM()
}

func (tvm *testVM) testLD(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.LD]
	argDir := arg{consts.TDirIdCode, 42}
	argReg := arg{consts.TRegIdCode, 2}

	Ld(tvm.p, argDir, argReg)
	if tvm.p.regs[argReg.val-1] != argDir.val {
		t.Errorf("direct arg")
	}

	memVal := -322
	argInd := arg{consts.TIndIdCode, consts.IdxMod + 2}
	tvm.vm.field.putInt32(argInd.val%consts.IdxMod, memVal)
	Ld(tvm.p, argInd, argReg)
	if tvm.p.regs[argReg.val-1] != memVal {
		t.Errorf("indirect arg")
	}

	Ld(tvm.p, arg{consts.TDirIdCode, 0}, argReg)
	if !tvm.p.carry {
		t.Errorf("carry shoul be 1")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testST(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.ST]
	argReg1 := arg{consts.TRegIdCode, 1}
	argReg2 := arg{consts.TRegIdCode, 2}

	memVal := 983
	tvm.p.storeReg(argReg1.val, memVal)
	St(tvm.p, argReg1, argReg2)
	if tvm.p.loadReg(argReg2.val) != memVal {
		t.Errorf("two registers")
	}

	memVal = -8733
	argReg := arg{consts.TRegIdCode, 1}
	argInd := arg{consts.TIndIdCode, len(tvm.vm.field.m) + consts.IdxMod}
	tvm.p.storeReg(argReg1.val, memVal)
	St(tvm.p, argReg, argInd)
	if tvm.vm.field.getInt32(argInd.val%consts.IdxMod) != memVal {
		t.Errorf("indirect")
	}

	*tvm = *newTestVM()
}

func (tvm *testVM) testAdd(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.ADD]
	argReg1 := arg{consts.TRegIdCode, 1}
	argReg2 := arg{consts.TRegIdCode, 2}
	argReg3 := arg{consts.TRegIdCode, 2}

	memVal1 := 983
	memVal2 := -62
	tvm.p.storeReg(argReg1.val, memVal1)
	tvm.p.storeReg(argReg2.val, memVal2)
	Add(tvm.p, argReg1, argReg2, argReg3)
	if tvm.p.loadReg(argReg3.val) != memVal1+memVal2 {
		t.Errorf("wrong sum")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testSub(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.SUB]
	argReg1 := arg{consts.TRegIdCode, 1}
	argReg2 := arg{consts.TRegIdCode, 2}
	argReg3 := arg{consts.TRegIdCode, 16}

	memVal1 := 983
	memVal2 := -62
	tvm.p.storeReg(argReg1.val, memVal1)
	tvm.p.storeReg(argReg2.val, memVal2)
	Sub(tvm.p, argReg1, argReg2, argReg3)
	if tvm.p.loadReg(argReg3.val) != memVal1-memVal2 {
		t.Errorf("wrong Sub")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testAnd(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.AND]
	argReg1 := arg{consts.TRegIdCode, 1}
	argReg2 := arg{consts.TRegIdCode, 2}
	argReg3 := arg{consts.TRegIdCode, 16}

	memVal1 := 983
	memVal2 := -62
	tvm.p.storeReg(argReg1.val, memVal1)
	tvm.p.storeReg(argReg2.val, memVal2)
	And(tvm.p, argReg1, argReg2, argReg3)
	if tvm.p.loadReg(argReg3.val) != memVal1&memVal2 {
		t.Errorf("wrong And")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testOr(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.OR]
	argReg1 := arg{consts.TRegIdCode, 1}
	argReg2 := arg{consts.TRegIdCode, 2}
	argReg3 := arg{consts.TRegIdCode, 16}

	memVal1 := 983
	memVal2 := -62
	tvm.p.storeReg(argReg1.val, memVal1)
	tvm.p.storeReg(argReg2.val, memVal2)
	Or(tvm.p, argReg1, argReg2, argReg3)
	if tvm.p.loadReg(argReg3.val) != memVal1|memVal2 {
		t.Errorf("wrong Or")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testXor(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.XOR]
	argReg1 := arg{consts.TRegIdCode, 1}
	argReg2 := arg{consts.TRegIdCode, 2}
	argReg3 := arg{consts.TRegIdCode, 16}

	memVal1 := 983
	memVal2 := -62
	tvm.p.storeReg(argReg1.val, memVal1)
	tvm.p.storeReg(argReg2.val, memVal2)
	Xor(tvm.p, argReg1, argReg2, argReg3)
	if tvm.p.loadReg(argReg3.val) != memVal1^memVal2 {
		t.Errorf("wrong Or")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testZjmp(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.ZJMP]
	// with carry

	currPC := 10
	jmpLen := consts.IdxMod + 10
	tvm.p.carry = true
	tvm.p.pc = currPC
	argDir1 := arg{consts.TDirIdCode, jmpLen}
	Zjmp(tvm.p, argDir1)
	if tvm.p.pc != currPC+jmpLen%consts.IdxMod {
		t.Errorf("wrong zjmp with carry")
	}

	// w/o carry
	tvm.p.carry = false
	currPC = tvm.p.pc
	Zjmp(tvm.p, argDir1)
	if tvm.p.pc != currPC {
		t.Errorf("wrong zjmp w/o carry")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testLdi(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.LDI]
	// ind dir reg

	argInd1 := arg{consts.TIndIdCode, 6}
	argReg3 := arg{consts.TRegIdCode, 16}

	i1 := consts.IdxMod + 1
	tvm.vm.field.putInt32(argInd1.val, i1)
	i2 := 33
	argDir2 := arg{consts.TDirIdCode, i2}
	val := -24
	tvm.vm.field.putInt32(tvm.p.pc+(i1+i2)%consts.IdxMod, val)
	Ldi(tvm.p, argInd1, argDir2, argReg3)
	if tvm.p.loadReg(argReg3.val) != val {
		t.Errorf("ldi: ind ind reg error")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testSti(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.STI]
	// ind dir reg

	argReg1 := arg{consts.TRegIdCode, 16}
	argInd2 := arg{consts.TIndIdCode, 34}
	argReg3 := arg{consts.TRegIdCode, 3}

	val := 8984
	tvm.p.storeReg(argReg1.val, val)
	i1 := consts.IdxMod + 1
	tvm.vm.field.putInt32(argInd2.val, i1)
	i2 := 33
	tvm.p.storeReg(argReg3.val, i2)
	Sti(tvm.p, argReg1, argInd2, argReg3)
	if tvm.vm.field.getInt32(tvm.p.pc+(i1+i2)%consts.IdxMod) != val {
		t.Errorf("sti error")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testFork(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.FORK]

	argDir1 := arg{consts.TDirIdCode, consts.IdxMod + 8}

	tvm.p.regs[0] = 4
	tvm.p.regs[15] = -34
	Fork(tvm.p, argDir1)
	if tvm.vm.procs.l.pc != argDir1.val%consts.IdxMod {
		t.Errorf("wrong forked pc")
		return
	}
	if !reflect.DeepEqual(tvm.vm.procs.l.regs, tvm.p.regs) {
		t.Errorf("not equal")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testLLD(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.LLD]
	argDir := arg{consts.TDirIdCode, 42}
	argReg := arg{consts.TRegIdCode, 2}

	Ld(tvm.p, argDir, argReg)
	if tvm.p.regs[argReg.val-1] != argDir.val {
		t.Errorf("direct arg")
	}

	memVal := -322
	argInd := arg{consts.TIndIdCode, consts.IdxMod + 2}
	tvm.vm.field.putInt32(argInd.val, memVal)
	Lld(tvm.p, argInd, argReg)
	if tvm.p.regs[argReg.val-1] != memVal {
		t.Errorf("indirect arg")
	}

	Ld(tvm.p, arg{consts.TDirIdCode, 0}, argReg)
	if !tvm.p.carry {
		t.Errorf("carry shoul be 1")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testLldi(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.LLDI]
	// ind dir reg

	argInd1 := arg{consts.TIndIdCode, 6}
	argReg3 := arg{consts.TRegIdCode, 16}

	i1 := consts.IdxMod + 1
	tvm.vm.field.putInt32(argInd1.val, i1)
	i2 := 33
	argDir2 := arg{consts.TDirIdCode, i2}
	val := -24
	tvm.vm.field.putInt32(tvm.p.pc+(i1+i2), val)
	Lldi(tvm.p, argInd1, argDir2, argReg3)
	if tvm.p.loadReg(argReg3.val) != val {
		t.Errorf("ldi: ind ind reg error")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testLfork(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.LFORK]

	argDir1 := arg{consts.TDirIdCode, consts.IdxMod + 8}

	tvm.p.regs[0] = 4
	tvm.p.regs[15] = -34
	Lfork(tvm.p, argDir1)
	if tvm.vm.procs.l.pc != argDir1.val {
		t.Errorf("wrong forked pc")
		return
	}
	if !reflect.DeepEqual(tvm.vm.procs.l.regs, tvm.p.regs) {
		t.Errorf("not equal")
	}
	*tvm = *newTestVM()
}
