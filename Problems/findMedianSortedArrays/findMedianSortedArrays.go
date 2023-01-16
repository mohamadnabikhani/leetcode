package findmediansortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	isEven := false
	medianIndex := 0
	// medianEvenInex := 0

	if (n+m)%2 == 0 {
		isEven = true
		medianIndex = (n + m) / 2
		// medianEvenInex = medianIndex + 1
	} else {
		medianIndex = ((n + m) / 2) + 1
	}

	result := 0.00
	curr := 0
	m1 := 0
	n1 := 0
	if (m + n) == 1 {
		if m == 0 {
			return float64(nums2[0])
		} else {
			return float64(nums1[0])
		}
	}
	for i := 0; i <= medianIndex; i++ {

		if m1 > len(nums1)-1 {
			curr = nums2[n1]
			n1 += 1
		} else if n1 > len(nums2)-1 {
			curr = nums1[m1]
			m1 += 1
		} else if nums1[m1] < nums2[n1] {
			curr = nums1[m1]
			m1 += 1

		} else {
			curr = nums2[n1]
			n1 += 1
		}
		if isEven && (i == (medianIndex - 1)) {
			result = float64(curr)
		} else if isEven && (i == (medianIndex)) {
			result = (result + float64(curr)) / 2
		} else if !isEven && (i == (medianIndex - 1)) {
			result = float64(curr)
		}
	}
	return result
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}
