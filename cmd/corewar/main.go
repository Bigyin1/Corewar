package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Bigyin1/Corewar/pkg/config"
	"github.com/Bigyin1/Corewar/pkg/github.com/Bigyin1/Corewar"
)

func readCfg(c *config.Config) {
	flag.BoolVar(&c.Log, "log", false, "log execution history to stdout")
	flag.IntVar(&c.Dump, "dump", -1,
		"dump game map state during provided cycle to stdout, then stop execution")
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
		c.Players = append(c.Players, config.PlayerData{Data: f})
	}
}

func main() {
	var cfg config.Config

	readCfg(&cfg)
	vm := github.com / Bigyin1 / Corewar.NewVM(&cfg)
	if err := vm.Init(cfg.Players...); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	vm.IntroducePlayers(os.Stdout)
	for !vm.Cycle() {
	}
	if vm.IsEnded() {
		winnerName := vm.GetWinner()
		fmt.Printf("Player %s won!\n", winnerName)
	}
}
