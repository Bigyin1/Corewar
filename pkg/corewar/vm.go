package corewar

import "corewar/pkg/consts"

type procList struct {
	l *proc
}

func (pl *procList) Put(np *proc) {
	if pl.l == nil {
		pl.l = np
		return
	}
	np.next = pl.l
	pl.l = np
}

func (pl *procList) Delete(id int) {
	var prev *proc
	currProc := pl.l
	for currProc != nil {
		if currProc.id != id {
			prev = currProc
			currProc = prev.next
			continue
		}
		if prev == nil {
			pl.l = currProc.next
			return
		}
		prev.next = currProc.next
		return
	}
}

func (pl *procList) IsEmpty() bool {
	if pl.l == nil {
		return true
	}
	return false
}

func (pl procList) exec(f func(p *proc)) {
	currProc := pl.l
	for currProc != nil {
		f(currProc)
		currProc = currProc.next
	}
}

type VM struct {
	players      []player
	lastAlive    *player
	procs        procList
	field        *field
	currCycle    int
	liveOps      int
	cyclesToDie  int
	eqInARow     int // for ho many checks, cyclesToDie is equal
	checksPassed int
	dump         bool
}

func (vm *VM) check() {
	if vm.cyclesToDie > 0 && vm.currCycle%vm.cyclesToDie != 0 {
		return
	}
	var toDel []int
	vm.procs.exec(func(p *proc) {
		if vm.currCycle-p.liveCycle >= vm.cyclesToDie {
			toDel = append(toDel, p.id)
			return
		}
	})
	for _, did := range toDel {
		vm.procs.Delete(did)
	}

	vm.checksPassed++
	if vm.liveOps >= consts.NbrLive {
		vm.cyclesToDie -= consts.CycleData
		vm.eqInARow = 0
		return
	}
	vm.eqInARow++
	if vm.eqInARow == consts.MaxChecks {
		vm.cyclesToDie -= consts.CycleData
		vm.eqInARow = 0
	}
}

func (vm *VM) Cycle() (isEnd bool) {
	vm.procs.exec((*proc).Cycle)
	vm.check()
	if vm.procs.IsEmpty() {
		isEnd = true
		return
	}
	vm.currCycle++
	return
}
