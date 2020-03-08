package tokenizer

import (
	"fmt"
	"strconv"
	"strings"
)

func (t *Tokenizer) isLabel(currStr string) (string, bool) {
	label := strings.Builder{}
	hasTerminate := false
	for idx := range currStr {
		if currStr[idx] == ':' {
			label.WriteByte(':')
			hasTerminate = true
			break
		}
		if !strings.Contains(labelChars, string(currStr[idx])) {
			return "", false
		}
		label.WriteByte(currStr[idx])
	}
	if label.Len() > 1 && hasTerminate {
		return label.String(), true
	}
	return "", false
}

func (t *Tokenizer) isString(currStr string) (string, bool) {
	if currStr[0] != '"' {
		return "", false
	}
	currStr = currStr[1:]
	res := strings.Builder{}
	res.WriteByte('"')
	for idx := range currStr {
		if currStr[idx] == '"' {
			res.WriteByte('"')
			return res.String(), true
		}
		res.WriteByte(currStr[idx])
	}
	return "", false
}

func (t *Tokenizer) isInstruction(currStr string) (string, bool) {
	currInstr := ""
	for instrName := range Instructions {
		if strings.HasPrefix(currStr, instrName) && len(instrName) > len(currInstr) {
			currInstr = instrName
		}
	}
	if currInstr == "" {
		return "", false
	}
	return currInstr, true
}

func (t *Tokenizer) isRegister(currStr string) (RegisterTokenVal, bool) {
	res := strings.Builder{}
	if currStr[0] != 'r' {
		return "", false
	}
	res.WriteByte('r')
	currStr = currStr[1:]
	num := strings.Builder{}
	for i := range currStr {
		if currStr[i] >= '0' && currStr[i] <= '9' {
			num.WriteByte(currStr[i])
			continue
		}
		break
	}
	n, err := strconv.Atoi(num.String())
	if err != nil {
		return "", false
	}
	if n > RegNumber || n == 0 {
		return "", false
	}
	res.WriteString(num.String())
	return RegisterTokenVal(res.String()), true
}

func (t *Tokenizer) isSeparator(currStr string) (string, bool) {
	if strings.HasPrefix(currStr, SeparatorSymbol) {
		return SeparatorSymbol, true
	}
	return "", false
}

func (t *Tokenizer) isDirect(currStr string) (string, bool) {
	res := strings.Builder{}
	if currStr[0] != '%' {
		return "", false
	}
	res.WriteByte('%')
	val, err := parserIntFromStr(currStr[1:])
	if err != nil {
		return "", false
	}
	_, _ = res.WriteString(strconv.Itoa(val))
	return res.String(), true
}

func (t *Tokenizer) isInDirect(currStr string) (InDirectTokenVal, bool) {
	res := strings.Builder{}
	val, err := parserIntFromStr(currStr)
	if err != nil {
		return "", false
	}
	_, _ = res.WriteString(strconv.Itoa(val))
	return InDirectTokenVal(res.String()), true
}

func (t *Tokenizer) isDirectLabel(currStr string) (string, bool) {
	if len(currStr) <= 2 {
		return "", false
	}
	res := strings.Builder{}
	if currStr[0] != '%' && currStr[1] != ':' {
		return "", false
	}
	res.WriteByte('%')
	res.WriteByte(':')
	labelStr := currStr[2:]
	for i := range labelStr {
		if !strings.Contains(labelChars, string(labelStr[i])) {
			if i == 0 {
				return "", false
			}
			break
		}
		res.WriteByte(labelStr[i])
	}

	return res.String(), true
}

func (t *Tokenizer) isInDirectLabel(currStr string) (string, bool) {
	if len(currStr) <= 1 {
		return "", false
	}
	res := strings.Builder{}
	if currStr[0] != ':' {
		return "", false
	}
	res.WriteByte(':')
	labelStr := currStr[1:]
	for i := range labelStr {
		if !strings.Contains(labelChars, string(labelStr[i])) {
			if i == 0 {
				return "", false
			}
			break
		}
		res.WriteByte(labelStr[i])
	}
	return res.String(), true
}

func (t *Tokenizer) getTokenType() (Token, int, error) {
	currStr := t.input[t.currIdx:]
	if strings.HasPrefix(currStr, nameHeader) {
		return Token{
			Typ:   Name,
			Value: nameHeader,
		}, len(nameHeader), nil
	}

	if strings.HasPrefix(currStr, commentHeader) {
		return Token{
			Typ:   Comment,
			Value: commentHeader,
		}, len(commentHeader), nil
	}
	if strings.HasPrefix(currStr, " ") || strings.HasPrefix(currStr, "\t") {
		return Token{
			Typ:   Space,
			Value: string(currStr[0]),
		}, 1, nil
	}
	if strings.HasPrefix(currStr, "\n") {
		return Token{
			Typ:   LineBreak,
			Value: BreakLineTokenVal("\n"),
		}, 1, nil
	}
	if strings.HasPrefix(currStr, "+") {
		return Token{
			Typ:   Sum,
			Value: "+",
		}, 1, nil
	}
	if strings.HasPrefix(currStr, "-") {
		return Token{
			Typ:   Sub,
			Value: "-",
		}, 1, nil
	}
	if label, ok := t.isLabel(currStr); ok {
		return Token{
			Typ:   Label,
			Value: label,
		}, len(label), nil
	}

	if str, ok := t.isString(currStr); ok {
		return Token{
			Typ:   Str,
			Value: str,
		}, len(str), nil
	}

	if i, ok := t.isInstruction(currStr); ok {
		return Token{
			Typ:   Instr,
			Value: i,
		}, len(i), nil
	}

	if i, ok := t.isRegister(currStr); ok {
		return Token{
			Typ:   Register,
			Value: i,
		}, len(i), nil
	}

	if i, ok := t.isSeparator(currStr); ok {
		return Token{
			Typ:   Separator,
			Value: i,
		}, len(i), nil
	}

	if direct, ok := t.isDirect(currStr); ok {
		return Token{
			Typ:   Direct,
			Value: direct,
		}, len(direct), nil
	}

	if direct, ok := t.isDirectLabel(currStr); ok {
		return Token{
			Typ:   DirectLabel,
			Value: direct,
		}, len(direct), nil
	}
	if inDirect, ok := t.isInDirect(currStr); ok {
		return Token{
			Typ:   Indirect,
			Value: inDirect,
		}, len(inDirect), nil
	}
	if inDirect, ok := t.isInDirectLabel(currStr); ok {
		return Token{
			Typ:   IndirectLabel,
			Value: inDirect,
		}, len(inDirect), nil
	}
	return Token{}, 0, fmt.Errorf("got unknown token at pos: %d", t.currIdx)

}
