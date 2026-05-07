/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	low, high := 1, n

	for low <= high {
		mid := low + (high-low)/2
		res := guess(mid)

		if res == 0 {
			return mid
		} else if res < 0 {
			// mid is too high → look lower
			high = mid - 1
		} else {
			// mid is too low → look higher
			low = mid + 1
		}
	}

	return -1 // just in case (should never reach here)
}