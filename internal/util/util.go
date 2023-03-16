package util

import (
	"io"
	"log"
	"strings"
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
