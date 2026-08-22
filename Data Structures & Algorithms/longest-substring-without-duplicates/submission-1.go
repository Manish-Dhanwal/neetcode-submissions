func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	maxlen := 0
	left := 0

	for right :=0; right < len(s); right++{
		seen[s[right]]++
		
		for seen[s[right]] > 1 {
			seen[s[left]]--
			left ++
		}
		

		window := right - left + 1
		if window > maxlen{
			maxlen = window
		}

	}
return maxlen

}