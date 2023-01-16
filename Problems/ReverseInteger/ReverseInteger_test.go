package reverseinteger

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
		s    int
		want int
	}{
		// {name: "test 1", s: 123, want: 321},
		{name: "test 2", s: -123, want: -321},
		// {name: "test 3", s: 120, want: 21},
		// {name: "test 3", s: "AB", numRows: 1, want: "AB"},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			result := reverse(tt.s)
			expected := tt.want
			assertCorrectMessage(t, tt.name, expected, result)
		})
	}
}
