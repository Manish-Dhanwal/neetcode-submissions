func groupAnagrams(strs []string) [][]string {
	
group := make(map[string][]string)

for i := 0; i <len(strs); i++{
	sorted := sortString(strs[i])
	
	group[sorted] = append(group[sorted], strs[i])
	}
	
	
	result := [][]string{}
	for _, k := range group{
		result = append(result, k)
	} 
	return result
}
func sortString(s string) string{
	  chars := []byte(s)

	  sort.Slice(chars, func (i, j int) bool {
		 return chars[i] < chars[j]
	  })
	return string(chars)
}