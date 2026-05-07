class Solution:
	def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
		count = 0
		maxFound = 0
		for i in range(len(nums)):
			if nums[i] != 0 and nums[i] == 1:
				count += 1
			else:
				count = 0
			maxFound = max(count, maxFound)
		return maxFound