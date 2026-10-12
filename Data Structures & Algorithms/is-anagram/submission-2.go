func isAnagram(s string, t string) bool {
	seen := make(map[byte]int)
	if len(s) != len(t){
		return false
	}
	for i, _ := range s{
		seen[s[i]]++
		seen[t[i]]--
	}
	for _, val := range seen{
		if val > 0{
			return false
		}
	}
	return true
}
