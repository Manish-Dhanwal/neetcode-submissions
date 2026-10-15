type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	str := ""
	for i :=0; i < len(strs);i++{
		str += strconv.Itoa(len(strs[i])) + "#" + strs[i]
	}
	return str
}

func (s *Solution) Decode(encoded string) []string {
	i := 0
	var result []string
	for i<len(encoded){
		j := i
		for encoded[j] != '#'{
			j++
		}
		length, _  := strconv.Atoi(encoded[i:j])
		word := encoded[j+1: j+1+length]
		result = append(result, word)
		i = j+1+length

	}
	return result
}
