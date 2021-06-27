package corewar

import (
	"os"
	"testing"

	"github.com/Bigyin1/Corewar/pkg/config"
	"github.com/Bigyin1/Corewar/pkg/testhelpers"
)

func TestVM_Start(t *testing.T) {
	tstDataDir := "./testdata/"
	testPlayers := []string{"test1.cor", "test2.cor"}

	d := make([]config.PlayerData, 0, len(testPlayers))

	for _, fn := range testPlayers {
		f, err := os.Open(tstDataDir + fn)
		if err != nil {
			t.FailNow()
		}
		d = append(d, config.PlayerData{Data: f})
	}
	vm := NewVM(&config.Config{})
	err := vm.Init(d...)
	if err != nil {
		t.Errorf("got error: %s", err.Error())
		return
	}

	if !testhelpers.Equal(t, len(vm.players), len(testPlayers),
		"got wrong players number") {
		return
	}
}
