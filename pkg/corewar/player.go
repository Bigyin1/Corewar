package corewar

import "io"

type PlayerData struct {
	CustomID int
	Data     io.ReadCloser
}

type player struct {
	id      int
	name    string
	comment string
	size    int

	code []byte
}
