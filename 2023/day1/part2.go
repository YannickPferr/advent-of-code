package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func AnswerPart2(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += getCalibrationValuePart2(line)
	}

	return sum
}

func getCalibrationValuePart2(line string) int {
	l := 0
	r := len(line) - 1
	leftDigit := ""
	rightDigit := ""
	for l <= r {
		leftDigit = getDigit(line, l)
		rightDigit = getDigit(line, r)
		if leftDigit != "" && rightDigit != "" {
			break
		}

		if leftDigit == "" {
			l++
		}

		if rightDigit == "" {
			r--
		}
	}

	calibrationValue, err := strconv.Atoi(fmt.Sprintf("%s%s", leftDigit, rightDigit))
	if err != nil {
		fmt.Println("shouldn't happen")
	}
	return calibrationValue
}

func getDigit(line string, idx int) string {
	r := rune(line[idx])
	if unicode.IsDigit(r) {
		return string(line[idx])
	}

	digitWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, digitWord := range digitWords {
		if idx+len(digitWord) <= len(line) && line[idx:idx+len(digitWord)] == digitWord {
			return wordToDigit(digitWord)
		}
	}

	return ""
}

func wordToDigit(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return ""
	}
}
