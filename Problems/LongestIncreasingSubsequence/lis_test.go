package lis

import (
	"testing"
)

func TestTableTwoSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, name string, expected, result int) {
		t.Helper()
		if expected != result {
			t.Errorf("%q expected '%d' but got '%d'", name, expected, result)
		}
	}

	wordTestTable := []struct {
		name string
		nums []int
		want int
	}{
		{name: "test {10,9,2,5,3,7,101,18}", nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, want: 4},
		{name: "test {0,1,0,3,2,3}", nums: []int{0, 1, 0, 3, 2, 3}, want: 4},
		{name: "test {7,7,7,7,7,7,7}", nums: []int{7, 7, 7, 7, 7, 7, 7}, want: 1},
	}

	for _, tt := range wordTestTable {
		t.Run(tt.name, func(t *testing.T) {
			result := LengthOfLIS(tt.nums)
			expected := tt.want
			assertCorrectMessage(t, tt.name, expected, result)
		})
	}
}
