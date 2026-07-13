func isPalindrome(s string) bool {
cleaned := ""
    for _, ch := range s {
        if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
            cleaned += strings.ToLower(string(ch))
        }
    }
left := 0
right := len(cleaned)-1
for left < right{
	if cleaned[left] != cleaned[right]{
		return false
	}
	left++
	right--
}
return true
}
