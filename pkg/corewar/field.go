package corewar

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
