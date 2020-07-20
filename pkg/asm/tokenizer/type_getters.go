package tokenizer

import (
	"corewar/consts"
	"fmt"
	"strconv"
	"strings"
)

func (t *Tokenizer) isLabel(currStr string) (string, bool) {
	label := strings.Builder{}
	hasTerminate := false
	for idx := range currStr {
		if currStr[idx] == ':' {
			hasTerminate = true
			break
		}
		if !strings.Contains(consts.LabelChars, string(currStr[idx])) {
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

func (t *Tokenizer) isInstruction(currStr string) (consts.InstructionName, bool) {
	currInstr := consts.InstructionName("")
	for instrName := range consts.InstructionsConfig {
		if strings.HasPrefix(currStr, string(instrName)) && len(instrName) > len(currInstr) {
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
	if n > consts.RegNumber || n == 0 {
		return "", false
	}
	res.WriteString(num.String())
	return RegisterTokenVal(res.String()), true
}

func (t *Tokenizer) isSeparator(currStr string) (string, bool) {
	if strings.HasPrefix(currStr, consts.SeparatorSymbol) {
		return consts.SeparatorSymbol, true
	}
	return "", false
}

func (t *Tokenizer) isDirect(currStr string) (DirectTokenVal, bool) {
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
	return DirectTokenVal(res.String()), true
}

func (t *Tokenizer) isInDirect(currStr string) (IndirectTokenVal, bool) {
	res := strings.Builder{}
	val, err := parserIntFromStr(currStr)
	if err != nil {
		return "", false
	}
	_, _ = res.WriteString(strconv.Itoa(val))
	return IndirectTokenVal(res.String()), true
}

func (t *Tokenizer) isDirectLabel(currStr string) (DirectLabelTokenVal, bool) {
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
		if !strings.Contains(consts.LabelChars, string(labelStr[i])) {
			if i == 0 {
				return "", false
			}
			break
		}
		res.WriteByte(labelStr[i])
	}

	return DirectLabelTokenVal(res.String()), true
}

func (t *Tokenizer) isInDirectLabel(currStr string) (IndirectLabelTokenVal, bool) {
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
		if !strings.Contains(consts.LabelChars, string(labelStr[i])) {
			if i == 0 {
				return "", false
			}
			break
		}
		res.WriteByte(labelStr[i])
	}
	return IndirectLabelTokenVal(res.String()), true
}

func (t *Tokenizer) getTokenType() (Token, int, error) {
	currStr := t.input[t.currIdx:]
	if strings.HasPrefix(currStr, consts.NameHeader) {
		return Token{
			Type:  ChampName,
			Value: consts.NameHeader,
		}, len(consts.NameHeader), nil
	}

	if strings.HasPrefix(currStr, consts.CommentHeader) {
		return Token{
			Type:  ChampComment,
			Value: consts.CommentHeader,
		}, len(consts.CommentHeader), nil
	}
	if strings.HasPrefix(currStr, " ") || strings.HasPrefix(currStr, "\t") {
		return Token{
			Type:  Space,
			Value: string(currStr[0]),
		}, 1, nil
	}
	if strings.HasPrefix(currStr, "\n") {
		return Token{
			Type:  LineBreak,
			Value: "\n",
		}, 1, nil
	}
	if strings.HasPrefix(currStr, "+") {
		return Token{
			Type:  Sum,
			Value: "+",
		}, 1, nil
	}
	if strings.HasPrefix(currStr, "-") {
		return Token{
			Type:  Sub,
			Value: "-",
		}, 1, nil
	}
	if label, ok := t.isLabel(currStr); ok {
		return Token{
			Type:  Label,
			Value: label,
		}, len(label) + 1, nil
	}

	if str, ok := t.isString(currStr); ok {
		return Token{
			Type:  Str,
			Value: str,
		}, len(str), nil
	}

	if i, ok := t.isInstruction(currStr); ok {
		return Token{
			Type:  Instr,
			Value: i,
		}, len(i), nil
	}

	if i, ok := t.isRegister(currStr); ok {
		return Token{
			Type:  Register,
			Value: i,
		}, len(i), nil
	}

	if i, ok := t.isSeparator(currStr); ok {
		return Token{
			Type:  Separator,
			Value: i,
		}, len(i), nil
	}

	if direct, ok := t.isDirect(currStr); ok {
		return Token{
			Type:  Direct,
			Value: direct,
		}, len(direct), nil
	}

	if direct, ok := t.isDirectLabel(currStr); ok {
		return Token{
			Type:  DirectLabel,
			Value: direct,
		}, len(direct), nil
	}
	if inDirect, ok := t.isInDirect(currStr); ok {
		return Token{
			Type:  Indirect,
			Value: inDirect,
		}, len(inDirect), nil
	}
	if inDirect, ok := t.isInDirectLabel(currStr); ok {
		return Token{
			Type:  IndirectLabel,
			Value: inDirect,
		}, len(inDirect), nil
	}
	return Token{}, 0, fmt.Errorf("got unknown token at pos: %d", t.currIdx)

}
