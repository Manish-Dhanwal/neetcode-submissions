func longestConsecutive(nums []int) int {

longest := 0

for _, num := range nums{
	current := num
	lenght := 1

	 for contain(nums, current+1){
		current++
		lenght++
	 }

	if longest < lenght{
		longest = lenght
	}
}

return longest
}


func contain(nums []int, target int)bool{
	for _, num := range nums{
		if num == target{
			return true
		}
	}
	return false
}