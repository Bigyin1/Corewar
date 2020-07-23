package main

import (
	"corewar/pkg/corewar"
	"flag"
	"fmt"
	"os"
)

func main() {
	var pData []corewar.PlayerData

	flag.Parse()
	for _, fName := range flag.Args() {
		f, err := os.Open(fName)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to open %s\n", fName)
			os.Exit(1)
		}

		pData = append(pData, corewar.PlayerData{Data: f})
	}

	err := corewar.NewVM().Start(pData...)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

}
