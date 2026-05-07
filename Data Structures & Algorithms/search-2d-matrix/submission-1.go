func searchMatrix(matrix [][]int, target int) bool {
	totalRows := len(matrix)
	totalCols := len(matrix[0])

	topRow := 0
	bottomRow := totalRows - 1

	// do a binary search to find the row where the target might be
	for topRow <= bottomRow {
		midRow := (topRow + bottomRow) / 2

		if target > matrix[midRow][totalCols-1] {
			topRow = midRow + 1
		} else if target < matrix[midRow][0] {
			bottomRow = midRow - 1
		} else {
			// break out of the loop if we find a row where the target might live
			break
		}
	}

	if topRow > bottomRow {
		return false
	}

	// binary search within the found row that might have the target
	searchRow := (topRow + bottomRow) / 2
	left := 0
	right := totalCols - 1
	for left <= right {
		midCol := (left + right) / 2
		if target > matrix[searchRow][midCol] {
			left = midCol + 1
		} else if target < matrix[searchRow][midCol] {
			right = midCol - 1
		} else {
			return true
		}
	}

	return false
}
