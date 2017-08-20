package arrays_and_strings

func ZeroMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	columns := len(matrix[0])
	zeroRows := make([]bool, rows)
	zeroCols := make([]bool, columns)

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == 0 {
				zeroRows[i] = true
				zeroCols[j] = true
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if zeroRows[i] || zeroCols[j] {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}
