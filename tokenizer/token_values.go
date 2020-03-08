package tokenizer

import (
	"strconv"
)

type RegisterTokenVal string

type BreakLineTokenVal string

func (v BreakLineTokenVal) String() string {
	return "\\n"
}

type DirectTokenVal string

func (v DirectTokenVal) GetValue() int {
	r, _ := strconv.Atoi(string(v[1:]))
	return r
}

type InDirectTokenVal string

func (v InDirectTokenVal) GetValue() int {
	r, _ := strconv.Atoi(string(v))
	return r
}
