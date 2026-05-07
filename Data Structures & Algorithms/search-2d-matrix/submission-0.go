func searchMatrix(matrix [][]int, target int) bool {
	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] == target {
				return true
			}
		}
	}
	return false
}
