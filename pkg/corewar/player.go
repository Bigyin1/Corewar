package corewar

import "io"

type PlayerData struct {
	CustomID int
	Data     io.Reader
}

type player struct {
	id      int
	name    string
	comment string

	code []byte
}
