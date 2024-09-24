package stringutils

import (
	"strings"
	"unicode/utf8"
)

func Reversestring(str string) string {
	sb := strings.Builder{}
	for i := utf8.RuneCountInString(str); i > 0; i-- {
		sb.WriteByte(str[i-1])
	}
	return sb.String()
}
