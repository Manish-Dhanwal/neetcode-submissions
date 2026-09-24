func isAnagram(s string, t string) bool {

seen := make(map[string]int)

if len(s) != len(t){
	return false
}
for i:= 0; i< len(s);i++{
	seen[string(s[i])]++
	seen[string(t[i])]--
}
for _, val := range seen{
	if val > 0{
		return false
	}
}
return true

}

