package corewar

import (
	"corewar/pkg/consts"
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

}

type testVM struct {
	vm *VM
	p  *proc
}

func newTestVM() *testVM {
	playerID := 1
	pc := 0
	fieldSz := 32

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
	tvm.p.Live(arg)
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

	tvm.p.Ld(argDir, argReg)
	if tvm.p.regs[argReg.val-1] != argDir.val {
		t.Errorf("direct arg")
	}

	memVal := -322
	argInd := arg{consts.TIndIdCode, 2}
	tvm.vm.field.PutInt32(argInd.val, memVal)
	tvm.p.Ld(argInd, argReg)
	if tvm.p.regs[argReg.val-1] != memVal {
		t.Errorf("indirect arg")
	}

	tvm.p.Ld(arg{consts.TDirIdCode, 0}, argReg)
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
	tvm.p.storeInReg(argReg1.val, memVal)
	tvm.p.St(argReg1, argReg2)
	if tvm.p.loadFromReg(argReg2.val) != memVal {
		t.Errorf("two registers")
	}

	memVal = -8733
	argReg := arg{consts.TRegIdCode, 1}
	argInd := arg{consts.TIndIdCode, len(tvm.vm.field.m)}
	tvm.p.storeInReg(argReg1.val, memVal)
	tvm.p.St(argReg, argInd)
	if tvm.vm.field.GetInt32(argInd.val) != memVal {
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
	tvm.p.storeInReg(argReg1.val, memVal1)
	tvm.p.storeInReg(argReg2.val, memVal2)
	tvm.p.Add(argReg1, argReg2, argReg3)
	if tvm.p.loadFromReg(argReg3.val) != memVal1+memVal2 {
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
	tvm.p.storeInReg(argReg1.val, memVal1)
	tvm.p.storeInReg(argReg2.val, memVal2)
	tvm.p.Sub(argReg1, argReg2, argReg3)
	if tvm.p.loadFromReg(argReg3.val) != memVal1-memVal2 {
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
	tvm.p.storeInReg(argReg1.val, memVal1)
	tvm.p.storeInReg(argReg2.val, memVal2)
	tvm.p.And(argReg1, argReg2, argReg3)
	if tvm.p.loadFromReg(argReg3.val) != memVal1&memVal2 {
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
	tvm.p.storeInReg(argReg1.val, memVal1)
	tvm.p.storeInReg(argReg2.val, memVal2)
	tvm.p.Or(argReg1, argReg2, argReg3)
	if tvm.p.loadFromReg(argReg3.val) != memVal1|memVal2 {
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
	tvm.p.storeInReg(argReg1.val, memVal1)
	tvm.p.storeInReg(argReg2.val, memVal2)
	tvm.p.Xor(argReg1, argReg2, argReg3)
	if tvm.p.loadFromReg(argReg3.val) != memVal1^memVal2 {
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
	tvm.p.Zjmp(argDir1)
	if tvm.p.pc != currPC+jmpLen%consts.IdxMod {
		t.Errorf("wrong zjmp with carry")
	}

	// w/o carry
	tvm.p.carry = false
	currPC = tvm.p.pc
	tvm.p.Zjmp(argDir1)
	if tvm.p.pc != currPC {
		t.Errorf("wrong zjmp w/o carry")
	}
	*tvm = *newTestVM()
}

func (tvm *testVM) testLdi(t *testing.T) {
	tvm.p.cmdMeta = consts.InstructionsConfig[consts.LDI]
	// dir dir reg

	argReg1 := arg{consts.TRegIdCode, 1}
	argReg2 := arg{consts.TRegIdCode, 2}
	argReg3 := arg{consts.TRegIdCode, 16}

}
