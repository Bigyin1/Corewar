package tokenizer

import (
	"strconv"
	"strings"
)

func parserIntFromStr(str string) (int, error) {
	res := strings.Builder{}
	for i := range str {
		if str[i] >= '0' && str[i] <= '9' || str[i] == '-' {
			res.WriteByte(str[i])
			continue
		}
		break
	}
	i, err := strconv.Atoi(res.String())
	if err != nil {
		return 0, err
	}
	return i, nil
}
