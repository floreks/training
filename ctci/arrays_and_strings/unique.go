package arrays_and_strings

func IsUnique(s string) bool {
	charMap := make(map[rune]bool, len(s))
	for _, c := range s {
		if _, exists := charMap[c]; exists {
			return false
		}

		charMap[c] = true
	}

	return true
}

func IsUniqueRaw(s string) bool {
	const alphabetSize = 128
	unique := make([]bool, alphabetSize)

	if len(s) > alphabetSize {
		return false
	}

	for _, c := range s {
		if unique[c] {
			return false
		}

		unique[c] = true
	}

	return true
}
