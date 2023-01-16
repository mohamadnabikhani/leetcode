package longestsubstringwithoutrepeatingcharacters

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, name string, expected, result int) {
		t.Helper()
		if expected != result {
			t.Errorf("%q expected '%d' but got '%d'", name, expected, result)
		}
	}

	testTable := []struct {
		name string
		s    string
		want int
	}{
		{name: "test 1", s: "abcabcbb", want: 3},
		{name: "test 2", s: "bbbbb", want: 1},
		{name: "test 3", s: "pwwkew", want: 3},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			result := LengthOfLongestSubstring(tt.s)
			expected := tt.want
			assertCorrectMessage(t, tt.name, expected, result)
		})
	}
}
