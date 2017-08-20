package arrays_and_strings

import "strings"

func URLify(s string, l int) string {
	s = strings.Trim(s, " ")
	return strings.Replace(s, " ", "%20", -1)
}
