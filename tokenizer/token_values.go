package tokenizer

import (
	"strconv"
)

type RegisterTokenVal string

func (v RegisterTokenVal) GetValue() uint8 {
	rnum, _ := strconv.Atoi(string(v[1:]))
	return uint8(rnum)
}

type BreakLineTokenVal string

func (v BreakLineTokenVal) String() string {
	return "\\n"
}

type DirectTokenVal string

func (v DirectTokenVal) GetValue() int32 {
	r, _ := strconv.Atoi(string(v[1:]))
	return int32(r)
}

type IndirectTokenVal string

func (v IndirectTokenVal) GetValue() int16 {
	r, _ := strconv.Atoi(string(v))
	return int16(r)
}

type DirectLabelTokenVal string

func (v DirectLabelTokenVal) GetValue() string {
	return string(v[2:])
}

type IndirectLabelTokenVal string

func (v IndirectLabelTokenVal) GetValue() string {
	return string(v[1:])
}
