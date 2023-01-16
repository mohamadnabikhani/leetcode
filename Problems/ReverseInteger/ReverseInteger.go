package reverseinteger

import "math"

func reverse(x int) int {
	digitCounter := lenLoop(x)
	result := 0
	remainedDigit := x
	for i := 1; i <= digitCounter; i++ {
		digit := remainedDigit % 10
		remainedDigit /= 10
		tmpIncAmount := int(math.Pow(10, float64(digitCounter-i))) * digit
		if result > 0 {
			if result > math.MaxInt32-tmpIncAmount {
				return 0
			}
		} else {
			if tmpIncAmount < math.MinInt32-result {
				return 0
			}
		}
		result += int(math.Pow(10, float64(digitCounter-i))) * digit
	}
	return result
}

func Reverse(x int) int {
	return reverse(x)
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
