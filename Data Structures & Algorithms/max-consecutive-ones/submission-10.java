class Solution {
    public int findMaxConsecutiveOnes(int[] nums) {
       int count = 0;
	   int maxFound = 0;

	   for (int i = 0; i < nums.length; i++) {
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