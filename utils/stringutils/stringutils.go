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

func StringBuilder(str ...[]string) string {
	sb := strings.Builder{}
	for _, s := range str {
		for _, ss := range s {
			sb.WriteString(ss + "\n")
		}
	}
	return sb.String()
}
