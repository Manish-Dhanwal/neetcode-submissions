func twoSum(nums []int, target int) []int {
	exist := make(map[int]int)
	for i:=0; i < len(nums); i++{
		diff := target - nums[i]
		if _, ok := exist[diff]; ok{
			return []int{exist[diff], i}
		}
	exist[nums[i]] = i
	}
	 return []int{}
}
