package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func AnswerPart1(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += getCalibrationValuePart1(line)
	}

	return sum
}

func getCalibrationValuePart1(line string) int {
	l := 0
	r := len(line) - 1
	for l <= r {
		left := rune(line[l])
		right := rune(line[r])
		if unicode.IsDigit(left) && unicode.IsDigit(right) {
			break
		}

		if !unicode.IsDigit(left) {
			l++
		}

		if !unicode.IsDigit(right) {
			r--
		}
	}

	calibrationValue, err := strconv.Atoi(fmt.Sprintf("%s%s", string(line[l]), string(line[r])))
	if err != nil {
		fmt.Println("shouldn't happen")
	}
	return calibrationValue
}
