func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)

	for i := 0; i < len(nums); i++{
		seen[nums[i]]++
	}
	pairs := []Pair{}
	
	for i, t := range seen{
		fmt.Println(i, t)
		pairs = append(pairs, Pair{
			num:  i,
			freq: t,
    })
	}
	
	sort.Slice(pairs, func(i, j int) bool {
    return pairs[i].freq > pairs[j].freq
})
	result := []int{}
	for i :=0; i < k; i++{
		result = append(result, pairs[i].num)
	}
	
	return result
}

type Pair struct{
	num int
	freq int
} 
