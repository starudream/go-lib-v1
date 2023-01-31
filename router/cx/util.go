package cx

import (
	"fmt"
	"unicode"
)

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func stringNotEmpty(vs ...string) string {
	for i := 0; i < len(vs); i++ {
		if vs[i] != "" {
			return vs[i]
		}
	}
	return ""
}

func format(s string, v ...any) string {
	if len(v) == 0 {
		return s
	}
	return fmt.Sprintf(s, v...)
}
