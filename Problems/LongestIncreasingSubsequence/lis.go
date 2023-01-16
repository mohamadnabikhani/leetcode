package lis

import "errors"

// func LengthOfLIS(nums []int) int {
// 	return lengthOfLIS(nums)
// }

// func lengthOfLIS(nums []int) int {
// 	subSlices := make([][]int, len(nums))
// 	subSlices[0] = make([]int, len(nums))
// 	for i, num := range(nums) {
// 		if num >
// 	}
// }

var ErrNotFound = errors.New("target not found in array")

func UpperBound(array []int, target int) (int, error) {
	startIndex := 0
	endIndex := len(array) - 1
	var mid int
	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] > target {
			endIndex = mid - 1
		} else {
			startIndex = mid + 1
		}
	}

	//when target greater or equal than every element in array, startIndex will out of bounds
	if startIndex >= len(array) {
		return -1, ErrNotFound
	}
	return startIndex, nil
}
