package longestpalindrome

func RecursieLongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	} else {
		if checkPalindrome(s) {
			return s
		} else {
			if checkPalindrome(s[0 : len(s)-1]) {
				return s[0 : len(s)-1]
			} else if checkPalindrome(s[1:]) {
				return s[1:]
			} else {
				return RecursieLongestPalindrome(s[1 : len(s)-1])
			}
		}
	}
}

func nonRecursivelongestPalindrome(s string) string {
	maxSub := ""
	foundMaxSub := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			maxSub += string(s[j])
			if checkPalindrome(maxSub) && len(maxSub) > len(foundMaxSub) {
				foundMaxSub = maxSub
			}
		}
		maxSub = ""
	}
	return foundMaxSub
}

func longestPalindrome(s string) string {
	return nonRecursivelongestPalindrome((s))
}

func checkPalindrome(str string) (result bool) {
	for i := 0; i < (len(str)/2)+1; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}
