func majorityElement(nums []int) int {
    seen := make(map[int]int)
	for _, num := range nums{
		seen[num]++
	}
	count := 0
	val := 0
	for v, ch := range seen{
		if ch > count{
			count = ch
			val = v
		}
	}
	return val
}
