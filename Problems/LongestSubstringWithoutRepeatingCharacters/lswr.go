package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	foundResult := ""
	result := ""
OUTER:
	for pos, char := range s {
		result = string(char)
		if len(result) > len(foundResult) {
			foundResult = result
		}
		m[byte(char)] = true
		// INNER:
		for i := pos + 1; i < len(s); i++ {
			if m[byte(s[i])] {
				for k := range m {
					delete(m, k)
				}
				continue OUTER
			} else {
				m[byte(s[i])] = true
				result += string(s[i])
				if len(result) > len(foundResult) {
					foundResult = result
				}
			}
		}
	}
	return len(foundResult)
}

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}

// best solutions
func lengthOfLongestSubstringX(s string) int {

	leftIndex := 0
	record := make(map[byte]int)
	maxLength := 0
	for index, letter := range []byte(s) {
		if lastIndex, ok := record[letter]; ok {
			if lastIndex >= leftIndex {
				leftIndex = lastIndex + 1
				record[letter] = index
			} else {
				record[letter] = index
			}
		} else {
			record[letter] = index
		}
		if index-leftIndex+1 > maxLength {
			maxLength = index - leftIndex + 1
		}

	}
	return maxLength
}
