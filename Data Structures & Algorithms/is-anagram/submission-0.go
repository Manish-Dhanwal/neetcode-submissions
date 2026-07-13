func isAnagram(s string, t string) bool {
	m := make(map[string]int)
	if len(s) != len(t){
		return false
	}
	for i, _ := range s{
		m[string(s[i])]++
		m[string(t[i])]--
	}
	for i, _ := range m{
		if m[i] > 0{
			return false
		}
	}
	return true
}
