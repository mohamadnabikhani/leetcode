package cvp

// https://leetcode.com/problems/count-vowels-permutation
// "ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou", "ua"

// var permisionedLetters = [5]rune{'a', 'e', 'i', 'o', 'u'}

var mod = (int(1e9) + 7)

func CountVowelPermutation(n int) int {
	return countVowelPermutation(n)
}

func countVowelPermutation(n int) int {
	countVowel := CountVowel{0, 0, 0, 0, 0}
	recursiveCountVowelPermutation(n, &countVowel)
	return (countVowel.curA + countVowel.curE + countVowel.curI + countVowel.curO + countVowel.curU) % mod
}

// number of Vowel in the end of word with n letter
// if n=1, so there's one word for each letter in its end, and so on.
type CountVowel struct {
	curA int
	curE int
	curI int
	curO int
	curU int
}

func recursiveCountVowelPermutation(n int, countVowel *CountVowel) {
	if n == 0 {
		return
	} else if n == 1 {
		countVowel.curA = 1
		countVowel.curE = 1
		countVowel.curI = 1
		countVowel.curO = 1
		countVowel.curU = 1
	} else {
		recursiveCountVowelPermutation(n-1, countVowel)
		lastCurA := countVowel.curA
		lastCurE := countVowel.curE
		lastCurI := countVowel.curI
		lastCurO := countVowel.curO
		lastCurU := countVowel.curU
		countVowel.curA = (lastCurE + lastCurI + lastCurU) % mod
		countVowel.curE = (lastCurA + lastCurI) % mod
		countVowel.curI = (lastCurE + lastCurO) % mod
		countVowel.curO = lastCurI
		countVowel.curU = (lastCurI + lastCurO) % mod
	}
}
