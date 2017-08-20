package arrays_and_strings

// The assumption is that image is NxN table
func Rotate(image [][]int) [][]int {
	l := len(image) - 1 // We need to subtract 1 to match table indexing correctly.

	// Number of steps is equal to n^2/4 because we are swapping 4 elements at a time
	for i := 0; i < len(image)*len(image)/4; i++ {
		a, b, c, d := i/l, i%l+i/l, l-i/l, l-i%l-i/l
		// Do a four way swap of TopL, TopR, BotR, BotL elements
		image[a][b], image[b][c], image[c][d], image[d][a] = image[d][a], image[a][b], image[b][c], image[c][d]
	}

	return image
}
