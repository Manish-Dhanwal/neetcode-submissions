func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)

	for _, val := range nums{
		seen[val]++
	}
	var pair []Freq
	for key, fr := range seen{
		pair = append(pair, Freq{
			val : key,
			fre: fr,
		})
	}
	sort.Slice(pair, func(i, j int)bool{
		return pair[i].fre > pair[j].fre
	})
	var result []int
	for i:=0; i< k; i++{
		result = append(result, pair[i].val)
	}
	return result
}

type Freq struct{
	val int
	fre int
}