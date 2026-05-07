class Solution {
   public int[] getConcatenation(int[] nums) {
      int[] ans = new int[2 * nums.length];
		int i = 0;
		int j = 0;

		while (i <= ans.length - 1) {
			ans[i] = nums[j];
			j++;
			if (j == nums.length) {
				j = 0;
			}
			i++;
		}
		return ans;
   }
}