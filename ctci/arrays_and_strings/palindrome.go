package arrays_and_strings

import (
	"strings"
	"unicode"
)

func IsPalindromePermutation(s string) bool {
	s = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	if len(s) == 0 || len(s) == 1 {
		return true
	}

	parityMap := make(map[rune]bool, 0)
	for _, r := range s {
		parityMap[r] = !parityMap[r]
	}

	foundMid := false
	for _, p := range parityMap {
		if p && !foundMid {
			foundMid = true
			continue
		}

		if p {
			return false
		}
	}

	return true
}

func IsPalindromePermutationBit(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	s = strings.ToLower(s)
	vector := 0
	mask := 1
	for i := range s {
		if nr := getCharNumber(s[i]); nr > -1 {
			mask = 1 << uint(nr)
			if vector&mask == 0 {
				vector |= mask
			} else {
				vector &^= mask
			}
		}
	}

	return (vector & (vector - 1)) == 0
}

func getCharNumber(c byte) int {
	if c < 'a' || c > 'z' {
		return -1
	}

	return int(c - 'a')
}
