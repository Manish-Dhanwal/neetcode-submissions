func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

	for i, num := range nums{
		check := target - num
		val, ok := seen[check]
		if ok{
			return []int{val, i}
		}
		seen[num] = i
	}
	return []int{}
}
