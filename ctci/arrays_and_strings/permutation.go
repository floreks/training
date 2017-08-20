package arrays_and_strings

func IsPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	countMap := make(map[byte]int, len(s1))
	for i := range s1 {
		countMap[s1[i]]++
		countMap[s2[i]]--
	}

	for _, counter := range countMap {
		if counter != 0 {
			return false
		}
	}

	return true
}
