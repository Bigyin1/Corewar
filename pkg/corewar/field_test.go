package corewar

import (
	"bytes"
	"testing"
)

func TestVMField(t *testing.T) {
	sz := 16
	f := newField(sz)

	tests := []struct {
		idx    int
		val    []byte
		intVal int
	}{
		{idx: 0, val: []byte{0x0, 0x0, 0xff, 0xff}, intVal: -32},
		{idx: sz, val: []byte{0xff, 0xff}},
		{idx: sz + sz, val: []byte{0x0, 0xff, 0xff}},
		{idx: sz / 2, val: []byte{0xff, 0xff, 0x0}},
		{idx: sz + 1, val: []byte{0xff, 0xff, 0x0, 0xfe}},
	}

	for i := range tests {
		rb := make([]byte, len(tests[i].val))
		wb := tests[i].val

		f.storeAt(tests[i].idx, wb)
		f.loadFrom(tests[i].idx, rb)
		if !bytes.Equal(wb, rb) {
			t.Errorf("stored And loaded bytes are not equal")
			return
		}
		f.putInt32(tests[i].idx, tests[i].intVal)
		if f.getInt32(tests[i].idx) != tests[i].intVal {
			t.Errorf("stored And loaded int32 are not equal")
		}
	}
}
