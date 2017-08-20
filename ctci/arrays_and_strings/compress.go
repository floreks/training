package arrays_and_strings

import (
	"bytes"
	"strconv"
)

func Compress(s string) string {
	buffer := bytes.NewBufferString("")

	count := 1
	for i := 0; i < len(s); i++ {
		tmp := i
		for tmp < len(s)-1 && s[tmp] == s[tmp+1] {
			tmp++
			count++
		}

		buffer.WriteByte(s[i])
		if count > 0 {
			buffer.WriteString(strconv.Itoa(count))
			i += count - 1
			count = 1
		}
	}

	if buffer.Len() >= len(s) {
		return s
	}

	return buffer.String()
}
