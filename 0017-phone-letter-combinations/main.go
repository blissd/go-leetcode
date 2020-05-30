package main

import (
	"strings"
)

var digitToLetters = map[uint8]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "qprs",
	'8': "tuv",
	'9': "wxyz",
}

func removeOnes(r rune) rune {
	if r == '1' {
		return -1
	}
	return r
}

func letterCombinations(digits string) []string {

	digits = strings.Map(removeOnes, digits)
	if len(digits) == 0 {
		return []string{}
	}

	totalCombos := len(digitToLetters[digits[0]])

	for i := 1; i < len(digits); i++ {
		totalCombos *= len(digitToLetters[digits[i]])
	}

	combos := make([][]uint8, totalCombos, totalCombos)
	for i := range combos {
		combos[i] = make([]uint8, len(digits), len(digits))
	}

	// compute one _column_ of letters at a time
	repeat := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		letters := digitToLetters[digit]
		// j == row, k = index into letters, count = repetition step
		for j, k, count := 0, 0, 0; j < totalCombos; j, count = j+1, count+1 {
			if count == repeat {
				count = 0
				k++
				if k == len(letters) {
					k = 0
				}
			}
			combos[j][i] = letters[k]
		}
		repeat *= len(letters)
	}

	result := make([]string, totalCombos, totalCombos)
	for i, v := range combos {
		result[i] = string(v)
	}

	return result
}
