package longestpalindrome

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, name string, expected, result string) {
		t.Helper()
		if expected != result {
			t.Errorf("%q expected '%q' but got '%q'", name, expected, result)
		}
	}

	testTable := []struct {
		name string
		s    string
		want string
	}{
		{name: "test 1", s: "babad", want: "babs"},
		{name: "test 2", s: "cbbd", want: "bb"},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.s[len(tt.s)-1:])
			result := longestPalindrome(tt.s)
			expected := tt.want
			assertCorrectMessage(t, tt.name, expected, result)
		})
	}
}

// best soution
func longestPalindromeX(s string) string {
	if len(s) == 1 {
		return s
	}
	//fmt.Println(isLongest([]byte("bb")))
	sc := []byte(s)
	for i := len(s); i > 0; i-- {
		start := 0

		for start+i <= len(s) {
			longe := sc[start : start+i]
			//fmt.Println(i,string(longe))
			if isLongest(longe) {
				return string(longe)
			}
			start++
		}
	}
	return ""
}
func isLongest(sc []byte) bool {

	start := 0
	end := len(sc) - 1
	for start < end {
		if sc[start] != sc[end] {
			return false
		}
		start++
		end--
	}
	return true
}
