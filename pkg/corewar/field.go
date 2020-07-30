package corewar

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type field struct {
	m []byte
}

func newField(sz int) *field {
	return &field{
		m: make([]byte, sz, sz),
	}
}

func (f *field) putCodeAt(idx int, code []byte) {
	copy(f.m[idx:], code)
}

func (f *field) loadFrom(idx int, d []byte) {
	idx %= len(f.m)
	if idx < 0 {
		idx += len(f.m)
	}
	if idx+len(d) <= len(f.m) {
		copy(d, f.m[idx:])
		return
	}
	c := copy(d, f.m[idx:])
	copy(d[c:], f.m[0:])
}

func (f *field) storeAt(idx int, d []byte) {
	idx %= len(f.m)
	if idx < 0 {
		idx += len(f.m)
	}

	if idx+len(d) <= len(f.m) {
		copy(f.m[idx:], d)
		return
	}

	c := copy(f.m[idx:], d)
	copy(f.m[0:], d[c:])
}

func (f *field) getInt32(idx int) int {
	var res int32
	var buf [4]byte
	f.loadFrom(idx, buf[:])
	_ = binary.Read(bytes.NewReader(buf[:]), binary.BigEndian, &res)
	return int(res)
}

func (f *field) getInt16(idx int) int {
	var res int16
	var buf [2]byte
	f.loadFrom(idx, buf[:])
	_ = binary.Read(bytes.NewReader(buf[:]), binary.BigEndian, &res)
	return int(res)
}

func (f *field) getByte(idx int) byte {
	var buf [1]byte
	f.loadFrom(idx, buf[:])
	return buf[0]
}

func (f *field) putInt32(idx int, val int) {
	var buf = bytes.NewBuffer(make([]byte, 0, 4))
	_ = binary.Write(buf, binary.BigEndian, int32(val))
	f.storeAt(idx, buf.Bytes())
	return
}

func (f *field) dump(w io.Writer) {
	perLine := 32
	for i, v := range f.m {
		if i%perLine == 0 {
			_, _ = fmt.Fprintf(w, "\n%#.4x : ", i)
		}
		_, _ = fmt.Fprintf(w, "%.2x ", v)
		//if i == perLine-1 {
		//	_, _ = fmt.Fprintf(w, "\n")
		//}
	}
}
