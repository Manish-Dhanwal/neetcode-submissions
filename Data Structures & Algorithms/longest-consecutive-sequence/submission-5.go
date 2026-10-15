func longestConsecutive(nums []int) int {
	seen := make(map[int]bool)
	longest := 0
	for _, val := range nums{
		seen[val] = true
	}

	for  _, val := range nums{
		if seen[val-1]{
			continue
		}
		current := val
		lenght := 1
		for seen[current+1]{
			current++
			lenght++
		}
		if lenght > longest{
			longest = lenght
		}

	}
return longest

}
