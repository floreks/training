package arrays_and_strings

func IsRotation(s1, s2 string) bool {
	if len(s1) != len(s2) || len(s1) == 0 || len(s2) == 0 {
		return false
	}

	s1 = s1 + s1
	return isSubstring(s1, s2)
}

func isSubstring(s1, s2 string) bool {
	idx2 := 0
	for i := range s1 {
		if s1[i] == s2[idx2] {
			idx2++
		} else {
			idx2 = 0
		}

		if idx2 == len(s2) {
			return true
		}
	}

	return false
}
