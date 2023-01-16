package atn

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, name string, expected, result *ListNode) {
		t.Helper()
		if !compareList(expected, result) {
			t.Errorf("%q expected '%q' but got '%q'", name, printList(expected), printList(result))
		}
	}

	testTable := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		// {name: "test 1", l1: insertBatchToList() (&ListNode{3, nil}).insertToList(4).insertToList(2), l2: (&ListNode{4, nil}).insertToList(6).insertToList(5), want: (&ListNode{8, nil}).insertToList(0).insertToList(7)},
		{name: "test 1", l1: insertBatchToList(3, 4, 2), l2: insertBatchToList(4, 6, 5), want: insertBatchToList(8, 0, 7)},
		// {name: "test 2", l1: &ListNode{0, nil}, l2: &ListNode{0, nil}, want: &ListNode{0, nil}},
		{name: "test 2", l1: insertBatchToList(0), l2: insertBatchToList(0), want: insertBatchToList(0)},
		// {name: "test 3", l1: (&ListNode{9, nil}).insertToList(9).insertToList(9).insertToList(9).insertToList(9).insertToList(9).insertToList(9), l2: (&ListNode{9, nil}).insertToList(9).insertToList(9).insertToList(9), want: (&ListNode{1, nil}).insertToList(0).insertToList(0).insertToList(0).insertToList(9).insertToList(9).insertToList(9).insertToList(8)},
		{name: "test 3", l1: insertBatchToList(9, 9, 9, 9, 9, 9, 9), l2: insertBatchToList(9, 9, 9, 9), want: insertBatchToList(1, 0, 0, 0, 9, 9, 9, 8)},
		{name: "test 3", l1: insertBatchToList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}...), l2: insertBatchToList([]int{5, 6, 4}...), want: insertBatchToList([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}...)},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			result := AddTwoNumbers(tt.l1, tt.l2)
			expected := tt.want
			assertCorrectMessage(t, tt.name, expected, result)
		})
	}
}
