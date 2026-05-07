class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    findMaxConsecutiveOnes(nums) {
		let count = 0;
		let maxFound = 0;
		for (let i = 0; i < nums.length; i++) {
			if (nums[i] != 0 && nums[i] == 1) {
				count++;
			} else {
				count = 0;
			}
			maxFound = count > maxFound ? count : maxFound;
		}
		return maxFound;
	}
}
