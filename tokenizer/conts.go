package tokenizer

const nameHeader = ".name"
const commentHeader = ".comment"
const labelChars = "abcdefghijklmnopqrstuvwxyz_0123456789"
const RegNumber = 16
const SeparatorSymbol = ","

var Instructions = map[string]struct{}{
	"live":  {},
	"ld":    {},
	"st":    {},
	"add":   {},
	"sub":   {},
	"and":   {},
	"or":    {},
	"xor":   {},
	"zjmp":  {},
	"ldi":   {},
	"sti":   {},
	"fork":  {},
	"lld":   {},
	"lldi":  {},
	"lfork": {},
	"aff":   {},
}
