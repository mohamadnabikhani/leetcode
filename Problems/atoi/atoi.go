package atoi

import (
	"math"
)

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	isNegative := false
	isEncounterNumber := false
	isEncounterSign := false
	isEncounterNonZeroNumber := false
	numberString := ""
	for _, char := range s {
		if !isEncounterSign && !isEncounterNumber {
			if char == ' ' {
				continue
			} else if char == '-' {
				isNegative = true
				isEncounterSign = true
				continue
			} else if char == '+' {
				isNegative = false
				isEncounterSign = true
				continue
			} else if byte(char) >= 48 && byte(char) <= 57 {
				isEncounterNumber = true
				if byte(char) == 48 {

				} else {
					isEncounterNonZeroNumber = true
					numberString += string(char)
				}
				continue
			} else {
				break
			}
		} else if !isEncounterNonZeroNumber {
			if byte(char) >= 48 && byte(char) <= 57 {
				if byte(char) == 48 {
					continue
				} else {
					numberString += string(char)
					isEncounterNonZeroNumber = true
					continue
				}
			} else {
				break
			}
		} else {
			if byte(char) >= 48 && byte(char) <= 57 {
				numberString += string(char)
			} else {
				break
			}
		}

	}
	resultNumber := 0
	signNumber := 1
	if isNegative {
		signNumber = -1
	}
	for i := 0; i < len(numberString); i++ {
		powerNumber := i
		stringDigitIndex := len(numberString) - 1 - i
		digit := int(numberString[stringDigitIndex] - '0')
		tenCounter := int(math.Pow(10, float64(powerNumber)))
		tmpAddedNumber := signNumber * digit * tenCounter
		// checkOverFlow := signNumber * 1 * tenCounter
		if digit != 0 {
			if tmpAddedNumber > 0 {
				if resultNumber > math.MaxInt32-tmpAddedNumber {
					return math.MaxInt32
				}
			} else {
				if tmpAddedNumber < math.MinInt32-resultNumber {
					return math.MinInt32
				}
			}
		} else {
			checkOverFlow := signNumber * 1 * tenCounter * 10
			if checkOverFlow > 0 {
				if checkOverFlow > math.MaxInt32-resultNumber {
					return math.MaxInt32
				}
			} else {
				if checkOverFlow < math.MinInt32-resultNumber {
					return math.MinInt32
				}
			}
		}

		resultNumber += tmpAddedNumber
	}
	return resultNumber
}

func MyAtoi(s string) int {
	return myAtoi(s)
}
