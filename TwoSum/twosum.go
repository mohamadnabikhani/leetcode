package twosum

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	for index, num := range nums {
		if foundedIndex := foundTargetIndexInSlice(nums, index, target-num); foundedIndex >= 0 {
			return []int{index, foundedIndex}
		}
	}
	return nil
}

func foundTargetIndexInSlice(nums []int, srcIndex, target int) int {
	for targetIndex := srcIndex + 1; targetIndex < len(nums); targetIndex++ {
		if srcIndex != targetIndex && nums[targetIndex] == target {
			return targetIndex
		}
	}
	return -1
}
