package cvp

import (
	"testing"
)

func TestTableDrivenCountVowelPermutation(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, name string, expected, sum int) {
		t.Helper()
		if expected != sum {
			t.Errorf("%q expected '%d' but got '%d'", name, expected, sum)
		}
	}

	wordTestTable := []struct {
		name string
		n    int
		want int
	}{
		{name: "test 1", n: 1, want: 5},
		{name: "test 2", n: 2, want: 10},
		{name: "test 3", n: 3, want: 19},
		{name: "test 4", n: 4, want: 35},
		{name: "test 5", n: 5, want: 68},
		{name: "test 144", n: 144, want: 18208803},
	}

	for _, tt := range wordTestTable {
		t.Run(tt.name, func(t *testing.T) {
			result := CountVowelPermutation(tt.n)
			expected := tt.want
			assertCorrectMessage(t, tt.name, expected, result)
		})
	}
}
