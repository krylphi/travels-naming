package util

import (
	"io"
	"log"
	"strings"
	"unicode"
)

func Concat(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	sb := strings.Builder{}
	for _, s := range strs {
		sb.WriteString(s)
	}
	return sb.String()
}

func CloseOrLog(c io.Closer) {
	if c == nil {
		return
	}
	err := c.Close()
	if err != nil {
		log.Println("RevGeoCode failed to close response body")
	}
}

func Title(sb strings.Builder) string {
	sep := " "
	ss := strings.SplitN(sb.String(), sep, 2)
	r := []rune(ss[0])
	r[0] = unicode.ToUpper(r[0])
	sb.Reset()
	sb.WriteString(string(r))
	sb.WriteString(sep)
	sb.WriteString(ss[1])
	return sb.String()
}
