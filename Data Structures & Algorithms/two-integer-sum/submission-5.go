func twoSum(nums []int, target int) []int {
   seen := make(map[int]int)

   for i, num := range nums{
		val := target - num
		_, ok := seen[val]
		if ok{
			return []int{seen[val], i}
		}
		seen[num] = i
   }
   return []int{}
}
