package config

import (
	"io"
)

type PlayerData struct {
	CustomID int
	Data     io.ReadCloser
}

type Config struct {
	Dump    int
	Log     bool
	Players []PlayerData
}
