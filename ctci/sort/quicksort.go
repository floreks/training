package sort

func max(i,j int) int {
	if i>j {
		return i
	}

	return j
}

func min(i,j int) int {
	if i<j {
		return i
	}

	return j
}

func median(i, j, k int) int {
	return max(min(i,j), min(max(i,j),k))
}

func partition(arr []int, l, h int) int {
	// Point around which we pivot could be calculated as median of elements at l, h, and (l+h)/2 indexes
	x := median(arr[l], arr[h], arr[(l+h)/2])
	i, j := l, h

	for true {
		// Look for value lower than x from the left
		for arr[i] < x {
			i++
		}

		// Look for value bigger than x from the right
		for arr[j] > x {
			j--
		}

		// If indexes overlap return end index
		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	return -1
}

func quicksort(arr []int, l, h int) {
	// Divide and sort arrays as long as l < h
	if l < h {
		p := partition(arr, l, h)
		// Sort left side of p
		quicksort(arr, l, p)
		// Sort right side of p
		quicksort(arr, p+1, h)
	}
}

func Quicksort(arr []int) {
	quicksort(arr, 0, len(arr)-1)
}
