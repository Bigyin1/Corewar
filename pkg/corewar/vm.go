package corewar

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

type VM struct {
	players      []player
	lastAlive    *player
	procs        procList
	field        *field
	cyclesPassed int
	liveOps      int
	cyclesToDie  int
	checksPassed int
	started      bool
}

func (vm *VM) Cycle() {
	if !vm.started {
		return
	}

}
