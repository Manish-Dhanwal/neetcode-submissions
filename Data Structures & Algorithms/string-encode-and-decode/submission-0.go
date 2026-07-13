type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	newStr := ""
	for _, s := range strs{
		newStr += strconv.Itoa(len(s))+ "#" + s
	}
	return newStr
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	i := 0

	for i < len(encoded){
		j := i
		for encoded[j] != '#'{
			j++
		}
		lenght, _ := strconv.Atoi(encoded[i:j])
		word := encoded[j+1 : j+1+lenght]
		result = append(result, word)
		i = j + 1 + lenght
	}
	return result
}
