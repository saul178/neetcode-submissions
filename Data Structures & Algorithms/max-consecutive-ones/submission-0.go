func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxFound := 0
	for i := range nums {
		if nums[i] != 0 && nums[i] == 1 {
			count++
		} else {
			count = 0
		}
		maxFound = max(count, maxFound)
	}
	return maxFound
	
}
