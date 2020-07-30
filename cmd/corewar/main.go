package main

import (
	"corewar/pkg/corewar"
	"flag"
	"fmt"
	"os"
	"strings"
)

type config struct {
	log     bool
	players []corewar.PlayerData
}

func (c *config) Read() {
	flag.BoolVar(&c.log, "log", false, "log execution history to stdout")
	flag.Parse()

	for _, fName := range flag.Args() {

		if !strings.HasSuffix(fName, ".cor") {
			_, _ = fmt.Fprintf(os.Stderr, "%s: only .cor files allowed\n", fName)
			os.Exit(1)
		}
		f, err := os.Open(fName)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to open %s\n", fName)
			os.Exit(1)
		}
		c.players = append(c.players, corewar.PlayerData{Data: f})
	}
}

func main() {
	var cfg config

	cfg.Read()
	vm := corewar.NewVM(cfg.log)
	if err := vm.Init(cfg.players...); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	vm.IntroducePlayers(os.Stdout)
	for !vm.Cycle() {
	}
	winnerName := vm.GetWinner()
	fmt.Printf("Player %s won!\n", winnerName)
}
