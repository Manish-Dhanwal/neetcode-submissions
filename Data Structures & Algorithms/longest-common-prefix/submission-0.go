func longestCommonPrefix(strs []string) string {
    if len(strs) == 0{
		return ""
	}
	prefix := strs[0]
	for i := 0; i < len(prefix); i++{
		for j := 1; j < len(strs); j++{
			if i >= len(strs[j]){
				return prefix[:i]
			}
			if strs[0][i] != strs[j][i]{
				return prefix[:i]
			}
		}
	}
	return prefix
}
