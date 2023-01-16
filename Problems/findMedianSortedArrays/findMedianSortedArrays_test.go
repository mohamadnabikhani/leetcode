package findmediansortedarrays

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, name string, expected, result float64) {
		t.Helper()
		if expected != result {
			t.Errorf("%q expected '%f' but got '%f'", name, expected, result)
		}
	}

	testTable := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{name: "test 1", nums1: []int{1, 3}, nums2: []int{2}, want: 2.00000},
		{name: "test 2", nums1: []int{1, 2}, nums2: []int{3, 4}, want: 2.50000},
		{name: "test 3", nums1: []int{}, nums2: []int{1}, want: 1.00000},
		{name: "test 4", nums1: []int{}, nums2: []int{1, 2}, want: 1.50000},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			result := FindMedianSortedArrays(tt.nums1, tt.nums2)
			expected := tt.want
			assertCorrectMessage(t, tt.name, expected, result)
		})
	}
}
