package corewar

import (
	"corewar/pkg/consts"
	"fmt"
	"io"
	"log"
	"os"
)

type procList struct {
	l   *proc
	lId int
}

func (pl *procList) Put(np *proc) {
	if pl.l == nil {
		pl.l = np
		return
	}
	pl.lId++
	np.id = pl.lId
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

func (pl procList) Exec(f func(p *proc)) {
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
	log          bool
	dump         bool
}

func (vm *VM) check() {
	if vm.cyclesToDie > 0 && vm.currCycle%vm.cyclesToDie != 0 {
		return
	}
	var toDel []int
	vm.procs.Exec(func(p *proc) {
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

func (vm *VM) IntroducePlayers(w io.Writer) {
	_, _ = fmt.Fprintf(w, "Introducing contestants...\n")
	for _, p := range vm.players {
		_, _ = fmt.Fprintf(w, "* Player %d, weighting %d bytes, %s, (%s)\n",
			p.id,
			p.size,
			p.name,
			p.comment)
	}
}

func (vm *VM) GetWinner() string {
	return vm.lastAlive.name
}

func (vm *VM) Cycle() (isEnd bool) {
	defer func() {
		if err := recover(); err != nil {
			vm.field.dump(os.Stderr)
			log.Fatalln(err)
		}
	}()

	vm.procs.Exec((*proc).Cycle)
	vm.check()
	if vm.procs.IsEmpty() {
		isEnd = true
		return
	}
	vm.currCycle++
	return
}
