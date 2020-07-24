package corewar

import (
	"bytes"
	"encoding/binary"
)

type field struct {
	m []byte
}

func newField(sz int) *field {
	return &field{
		m: make([]byte, sz, sz),
	}
}

func (f *field) PutCodeAt(idx int, code []byte) {
	copy(f.m[idx:], code)
}

func (f *field) LoadFrom(idx int, d []byte) {
	idx %= len(f.m)
	if idx+len(d) <= len(f.m) {
		copy(d, f.m[idx:])
		return
	}
	c := copy(d, f.m[idx:])
	copy(d[c:], f.m[0:])
}

func (f *field) StoreAt(idx int, d []byte) {
	idx %= len(f.m)

	if idx+len(d) <= len(f.m) {
		copy(f.m[idx:], d)
		return
	}

	c := copy(f.m[idx:], d)
	copy(f.m[0:], d[c:])
}

func (f *field) GetInt32(idx int) (v int32) {
	var buf [4]byte
	f.LoadFrom(idx, buf[:])
	_ = binary.Read(bytes.NewReader(buf[:]), binary.BigEndian, &v)
	return
}

func (f *field) PutInt32(idx int, val int32) {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], uint32(val))
	f.StoreAt(idx, buf[:])
	return
}
