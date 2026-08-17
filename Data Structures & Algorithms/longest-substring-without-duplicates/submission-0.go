func lengthOfLongestSubstring(s string) int {

	seen := make(map[byte]bool)
	left := 0
	lenght := 0

	for right := 0; right < len(s); right++{
		char := s[right]
		for seen[char]{
			delete(seen, s[left])
			left++
		}

		seen[char] = true

		checkLen := right - left + 1

		if checkLen > lenght{
			lenght = checkLen
		}
	}
return lenght

}
