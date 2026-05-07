func replaceElements(arr []int) []int {
	rightMax := -1
	for i := len(arr) - 1; i >= 0; i-- {
		newMax := max(rightMax, arr[i])
		arr[i] = rightMax
		rightMax = newMax
	}
	return arr
}
