func lengthOfLongestSubstring(s string) int {
	// have a set to show the unique chars shown in the string
	if s == "" {
        return 0
    }
    
    charSet := make(map[byte]bool)
	substringLen := math.MinInt
	leftPtr := 0

	for righPtr := 0; righPtr < len(s); righPtr++ {
		for charSet[s[righPtr]] {
			delete(charSet, s[leftPtr])
			leftPtr++
		}
		charSet[s[righPtr]] = true
		substringLen = max(substringLen, righPtr-leftPtr+1)
	}
	return substringLen
}
