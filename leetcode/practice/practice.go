package practice

// 804唯一福尔斯密码词
func uniqueMorseRepresentations(words []string) int {
	str := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---",
		"-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]int)
	for _, v := range words {
		s := ""
		for _, t := range v {
			s += str[t-'a']
		}
		m[s] = 1
	}
	return len(m)
}
