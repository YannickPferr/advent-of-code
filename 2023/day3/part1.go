package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func AnswerPart1(lines []string) int {
	sum := 0
	for i, line := range lines {
		seenDigitWithAdjacentSymbol := false
		sb := strings.Builder{}
		for j := range line {
			if unicode.IsDigit(rune(line[j])) {
				if hasAdjacentSymbol(lines, i, j) {
					seenDigitWithAdjacentSymbol = true
				}
				sb.WriteString(string(line[j]))
				continue
			}

			if seenDigitWithAdjacentSymbol {
				val, err := strconv.Atoi(sb.String())
				if err != nil {
					fmt.Println("shouldn't happen")
				}
				sum += val
			}
			sb.Reset()
			seenDigitWithAdjacentSymbol = false
		}

		// need to check at the end again in case last char was digit
		if seenDigitWithAdjacentSymbol {
			val, err := strconv.Atoi(sb.String())
			if err != nil {
				fmt.Println("shouldn't happen")
			}
			sum += val
		}
	}

	return sum
}

func hasAdjacentSymbol(lines []string, i int, j int) bool {
	if isSymbol(lines, i-1, j) {
		return true
	}

	if isSymbol(lines, i, j-1) {
		return true
	}

	if isSymbol(lines, i+1, j) {
		return true
	}

	if isSymbol(lines, i, j+1) {
		return true
	}

	if isSymbol(lines, i-1, j-1) {
		return true
	}

	if isSymbol(lines, i-1, j+1) {
		return true
	}

	if isSymbol(lines, i+1, j-1) {
		return true
	}

	if isSymbol(lines, i+1, j+1) {
		return true
	}

	return false
}

func isSymbol(lines []string, i, j int) bool {
	// out of bounds check
	if i < 0 || j < 0 || i >= len(lines) || j >= len(lines[0]) {
		return false
	}
	r := rune(lines[i][j])
	return !unicode.IsDigit(r) && r != '.'
}
