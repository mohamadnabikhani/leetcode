package zigzagconversion

import (
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
		name    string
		s       string
		numRows int
		want    string
	}{
		{name: "test 1", s: "PAYPALISHIRING", numRows: 3, want: "PAHNAPLSIIGYIR"},
		{name: "test 2", s: "PAYPALISHIRING", numRows: 4, want: "PINALSIGYAHRPI"},
		{name: "test 3", s: "A", numRows: 1, want: "A"},
		{name: "test 3", s: "AB", numRows: 1, want: "AB"},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			result := convert(tt.s, tt.numRows)
			expected := tt.want
			assertCorrectMessage(t, tt.name, expected, result)
		})
	}
}
