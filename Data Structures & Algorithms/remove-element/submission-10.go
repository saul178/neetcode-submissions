func removeElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			copy(nums[i:], nums[i+1:])
			nums[len(nums)-1] = 0
			nums = nums[:len(nums)-1]
		} else {
            i++
        }
	}
	return len(nums)
}
