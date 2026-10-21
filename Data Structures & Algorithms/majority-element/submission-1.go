func majorityElement(nums []int) int {
    // seen := make(map[int]int)
	// for _, num := range nums{
	// 	seen[num]++
	// }
	// count := 0
	// val := 0
	// for v, ch := range seen{
	// 	if ch > count{
	// 		count = ch
	// 		val = v
	// 	}
	// }
	// return val

	candidate := 0
	count := 0
	for _, num := range nums{
		if count == 0{
			candidate = num
		}
		if num == candidate{
			count++
		}else{
			count--
		}
	}
	return candidate
}
