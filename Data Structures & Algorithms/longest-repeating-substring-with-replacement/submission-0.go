func characterReplacement(s string, k int) int {
seen := make(map[byte]int)
left := 0
maxLen := 0
maxFrequency := 0

for right := 0; right < len(s); right++{
	char := s[right]
	seen[char]++

	
	if maxFrequency < seen[char] {
		maxFrequency = seen[char]
	}

	window := right - left + 1
	replacement := window - maxFrequency 
	if replacement > k{
		seen[s[left]]--
		left++
	}
	window = right - left + 1
	if window > maxLen{
		maxLen = window
	}
}
return maxLen
}
