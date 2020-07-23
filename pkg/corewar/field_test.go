package corewar

import (
	"bytes"
	"testing"
)

func TestVMField(t *testing.T) {
	sz := 16
	f := newField(sz)

	tests := []struct {
		idx int
		val []byte
	}{
		{idx: 0, val: []byte{0xff, 0xff, 0x0}},
		{idx: sz, val: []byte{0xff, 0xff, 0x0}},
		{idx: sz + sz, val: []byte{0xff, 0xff, 0x0}},
		{idx: sz / 2, val: []byte{0xff, 0xff, 0x0}},
		{idx: sz + 1, val: []byte{0xff, 0xff, 0x0, 0xfe}},
	}

	for i := range tests {
		rb := make([]byte, len(tests[i].val))
		wb := tests[i].val

		f.StoreAt(tests[i].idx, wb)
		f.LoadFrom(tests[i].idx, rb)
		if !bytes.Equal(wb, rb) {
			t.Errorf("stored and loaded bytes are not equal")
			return
		}
	}
}
