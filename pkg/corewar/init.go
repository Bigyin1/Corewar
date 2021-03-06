package corewar

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"sort"

	"github.com/Bigyin1/Corewar/pkg/config"
	"github.com/Bigyin1/Corewar/pkg/consts"
)

func NewVM(cfg *config.Config) *VM {

	return &VM{
		cyclesToDie: consts.CyclesToDie,
		log:         cfg.Log,
		dump:        cfg.Dump,
	}
}

func (vm *VM) Init(pd ...config.PlayerData) error {
	var pCount int

	pCount = len(pd)
	if pCount < 2 {
		return fmt.Errorf("got not enough players: %d", pCount)
	}
	if pCount > consts.MaxPlayers {
		return fmt.Errorf("too many players - %d, %d is max", pCount, consts.MaxPlayers)
	}
	if err := setupPlayersIDs(pd); err != nil {
		return err
	}
	if err := vm.loadPlayersMeta(pd); err != nil {
		return err
	}
	vm.lastAlive = &vm.players[len(vm.players)-1]
	vm.initProcs()
	return nil
}

func validateIDs(p []config.PlayerData) error {
	for i := 0; i < len(p)-1; i++ {
		if p[i].CustomID == p[i+1].CustomID && p[i].CustomID != 0 {
			return fmt.Errorf("duplicated id: %d", p[i].CustomID)
		}
	}
	return nil
}

func setupPlayersIDs(p []config.PlayerData) error {
	if err := validateIDs(p); err != nil {
		return err
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].CustomID < p[j].CustomID
	})
	idsList := make([]int, len(p))
	for i := 1; i <= len(p); i++ {
		idsList[i-1] = i
	}

	for i := range p {
		p[i].CustomID = idsList[i]
	}
	return nil
}

type playerHeader struct {
	Magic    [len(consts.MagicHeader)]byte
	Name     [consts.ProgNameLength]byte
	Null1    [len(consts.NullSeq)]byte
	CodeSize int32
	Comment  [consts.CommentLength]byte
	Null2    [len(consts.NullSeq)]byte
}

func parseHeader(d io.ReadCloser, id int) (player, error) {
	var p player
	var h playerHeader

	err := binary.Read(d, binary.BigEndian, &h)
	if err != nil {
		return player{}, fmt.Errorf("player %d, has invalid code header", id)
	}
	p.id = id
	if string(h.Magic[:]) != consts.MagicHeader {
		return player{}, fmt.Errorf("player %d, has invalid code magic header", id)
	}
	p.name = string(bytes.Split(h.Name[:], []byte{0})[0])
	if string(h.Null1[:]) != consts.NullSeq {
		return player{}, fmt.Errorf("player %d, has invalid code format", id)
	}
	if int(h.CodeSize) > consts.ChampMaxSize {
		return player{}, fmt.Errorf("player %d, has too long code (%d)", id, int(h.CodeSize))
	}
	p.size = int(h.CodeSize)
	p.comment = string(bytes.Split(h.Comment[:], []byte{0})[0])
	if string(h.Null2[:]) != consts.NullSeq {
		return player{}, fmt.Errorf("player %d, has invalid code format", id)
	}

	code, err := ioutil.ReadAll(d)
	if err != nil {
		return player{}, fmt.Errorf("player %d, has invalid code format", id)
	}
	if len(code) != int(h.CodeSize) {
		return player{}, fmt.Errorf("player %d, hs longer code, than was mentioned in header", id)
	}
	_ = d.Close()
	p.code = code
	return p, nil
}

func (vm *VM) loadPlayersMeta(pd []config.PlayerData) error {
	for i := range pd {
		p, err := parseHeader(pd[i].Data, pd[i].CustomID)
		if err != nil {
			return err
		}
		vm.players = append(vm.players, p)
	}
	return nil
}

func (vm *VM) initProcs() {
	vm.field = newField(consts.MemSize)

	var procID = 1
	var idx int
	var idxStep = consts.MemSize / len(vm.players)

	for i := range vm.players {
		vm.field.putCodeAt(idx, vm.players[i].code)
		vm.procs.Put(newProc(procID, vm.players[i].id, idx, vm))
		vm.procs.lId = procID

		vm.players[i].code = nil // deallocate
		procID++
		idx += idxStep
	}

}
