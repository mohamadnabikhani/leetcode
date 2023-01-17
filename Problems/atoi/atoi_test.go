package atoi

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
		{name: "test 1", s: "42", want: 42},
		{name: "test 2", s: "   -42", want: -42},
		{name: "test 3", s: "4193 with words", want: 4193},
		{name: "test 4", s: "+-12", want: 0},
		{name: "test 4", s: "20000000000000000000", want: 2147483647},
		{name: "test 4", s: "10000000000000003644", want: 2147483647},
		{name: "test 4", s: "-000000000000001", want: -1},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			result := myAtoi(tt.s)
			expected := tt.want
			assertCorrectMessage(t, tt.name, expected, result)
		})
	}
}
