func binarySearch(left int, right int, nums []int, target int) int {
    if left > right {
        return -1
    }

    mid := left + (right - left) / 2
    if nums[mid] == target {
        return mid
    }

    if nums[mid] < target {
        return binarySearch(mid+1, right, nums, target)
    }
    return binarySearch(left, mid-1, nums, target)

}

func search(nums []int, target int) int {
    return binarySearch(0, len(nums)-1, nums, target)
}
