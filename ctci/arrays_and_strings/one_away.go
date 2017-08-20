package arrays_and_strings

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func IsOneAway(s1, s2 string) bool {
	if abs(len(s1)-len(s2)) > 1 {
		return false
	}

	countMap := make(map[rune]int, 0)
	for _, r := range s1 {
		countMap[r]++
	}

	for _, r := range s2 {
		countMap[r]--
	}

	repeatCount := 0
	for _, count := range countMap {
		repeatCount += abs(count)

		if repeatCount > 2 {
			return false
		}
	}

	return true
}

func longerStringFirst(s1, s2 string) (string, string) {
	if len(s1) > len(s2) {
		return s1, s2
	}

	return s2, s1
}

func IsOneAwaySimple(s1, s2 string) bool {
	if abs(len(s1)-len(s2)) > 1 {
		return false
	}

	longer, shorter := longerStringFirst(s1, s2)
	foundDifference := false
	for idx1, idx2 := 0, 0; idx1 < len(longer) && idx2 < len(shorter); idx1, idx2 = idx1+1, idx2+1 {
		if longer[idx1] != shorter[idx2] {
			if foundDifference {
				return false
			}

			if len(s1) != len(s2) {
				if longer[idx1+1] != shorter[idx2] {
					return false
				}

				idx1++
			}

			foundDifference = true
		}
	}

	return true
}
