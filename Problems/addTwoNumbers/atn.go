package atn

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertToList(l *ListNode, d int) *ListNode {
	list := &ListNode{Val: d, Next: l}
	return list
}

func insertBatchToList(arr ...int) *ListNode {
	list := &ListNode{Val: arr[0], Next: nil}
	for i := 1; i < len(arr); i++ {
		list = insertToList(list, arr[i])
	}
	return list

}

func compareList(l1 *ListNode, l2 *ListNode) bool {
	compareTarget1 := l1
	compareTarget2 := l2
	for compareTarget1 != nil && compareTarget2 != nil {
		if compareTarget1.Val != compareTarget2.Val {
			return false
		}
		compareTarget1 = compareTarget1.Next
		compareTarget2 = compareTarget2.Next
	}
	if compareTarget1 == nil && compareTarget2 == nil {
		return true
	}
	return false
}

func printList(l *ListNode) string {
	target := l
	var result string
	for target != nil {
		result += fmt.Sprintf("-> %d ", target.Val)
		target = target.Next
	}
	return result
}

func lenLoop(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func AddTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0, nil}
	target1 := l1
	target2 := l2

	number1 := 0
	digitCounter1 := 0

	number2 := 0
	digitCounter2 := 0

	for target1 != nil {
		number1 += int(math.Pow(10, float64(digitCounter1))) * target1.Val
		target1 = target1.Next
		digitCounter1 += 1
	}

	for target2 != nil {
		number2 += int(math.Pow(10, float64(digitCounter2))) * target2.Val
		target2 = target2.Next
		digitCounter2 += 1
	}
	resultNumber := number1 + number2
	digitCounter := lenLoop(resultNumber)
	curVal := 0
	remainNumber := resultNumber
	firstInsert := true
	for digitCounter > 0 {
		digitCounter -= 1

		curVal = remainNumber / (int(math.Pow(10, float64(digitCounter))))
		if curVal != 0 {
			remainNumber %= (curVal * (int(math.Pow(10, float64(digitCounter)))))
		}
		if firstInsert {
			result.Val = curVal
			firstInsert = false
		} else {
			result = insertToList(result, curVal)
		}
	}

	return result
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}
		if tmp.Val > 9 {
			tmp.Val -= 10
			tmp.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next
	}
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}
