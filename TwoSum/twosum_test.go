package twosum

import (
	"reflect"
	"testing"
)

func TestTableTwoSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, name string, expected, sum []int) {
		t.Helper()
		if !reflect.DeepEqual(expected, sum) {
			t.Fatalf("%q: wanted %v, got %v", name, expected, sum)
		}
	}

	wordTestTable := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "test {2, 7, 11, 15} 9", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{name: "test {3, 2, 4} 6", nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{name: "test {3, 3} 6", nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}

	for _, tt := range wordTestTable {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoSum(tt.nums, tt.target)
			expected := tt.want
			assertCorrectMessage(t, tt.name, expected, result)
		})
	}
}
