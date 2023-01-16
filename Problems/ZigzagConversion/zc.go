package zigzagconversion

func convert(s string, numRows int) string {
	result := make([]rune, len(s))
	var lastIndexInRow int
	resultIndex := 0
	if numRows == 1 {
		return s
	}
	for i := 0; i < numRows; i++ {

		indexFactors := []int{2 * (numRows - 1 - i), 2 * (i)}
		j := 0
		lastIndexInRow = i
		for lastIndexInRow < len(s) && resultIndex < len(s) {
			result[resultIndex] = rune(s[lastIndexInRow])
			for (indexFactors[0] != indexFactors[1]) && indexFactors[j%2] == 0 {
				j++
			}
			lastIndexInRow += indexFactors[j%2]
			resultIndex++
			j++
		}
	}
	return string(result)
}

func Convert(s string, numRows int) string {
	return convert(s, numRows)
}
