func checkInclusion(s1 string, s2 string) bool {
	left := 0
	s1Count := make(map[byte]int)

	for i := 0; i < len(s1); i++{
		s1Count[s1[i]]++
	}

	window := make(map[byte]int)
	for right := 0; right < len(s2); right++{
		window[s2[right]]++

		windowSize := right - left + 1
		if windowSize > len(s1){
			window[s2[left]]--
			left++
		}
		match := true
		windowSize = right - left + 1
		if windowSize == len(s1){
			for char, count := range s1Count{
				if window[char] != count{
					match = false
					break
				}
			}
			if match{
			return true
		}
		}
		
	}
return false
}
