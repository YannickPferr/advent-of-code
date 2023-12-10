package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func AnswerPart2(lines []string) int {
	sum := 0
	for i, line := range lines {
		for j := range line {
			if rune(line[j]) == '*' {
				adjacentNumbers := getAdjacentNumbers(lines, i, j)
				if len(adjacentNumbers) == 2 {
					sum += adjacentNumbers[0] * adjacentNumbers[1]
				}
			}
		}
	}

	return sum
}

func getAdjacentNumbers(lines []string, i int, j int) []int {
	var adjacentNumbers []int
	visited := make([][]bool, len(lines))
	for idx, line := range lines {
		visited[idx] = make([]bool, len(line))
	}

	if num, err := getNumber(lines, i-1, j, visited); err == nil {
		adjacentNumbers = append(adjacentNumbers, num)
	}

	if num, err := getNumber(lines, i, j-1, visited); err == nil {
		adjacentNumbers = append(adjacentNumbers, num)
	}

	if num, err := getNumber(lines, i+1, j, visited); err == nil {
		adjacentNumbers = append(adjacentNumbers, num)
	}

	if num, err := getNumber(lines, i, j+1, visited); err == nil {
		adjacentNumbers = append(adjacentNumbers, num)
	}

	if num, err := getNumber(lines, i-1, j-1, visited); err == nil {
		adjacentNumbers = append(adjacentNumbers, num)
	}

	if num, err := getNumber(lines, i-1, j+1, visited); err == nil {
		adjacentNumbers = append(adjacentNumbers, num)
	}

	if num, err := getNumber(lines, i+1, j-1, visited); err == nil {
		adjacentNumbers = append(adjacentNumbers, num)
	}

	if num, err := getNumber(lines, i+1, j+1, visited); err == nil {
		adjacentNumbers = append(adjacentNumbers, num)
	}

	return adjacentNumbers
}

func getNumber(lines []string, i, j int, visited [][]bool) (int, error) {
	// out of bounds check
	if i < 0 || j < 0 || i >= len(lines) || j >= len(lines[0]) {
		return -1, errors.New("index out of bounds")
	}
	// check if digit
	if !unicode.IsDigit(rune(lines[i][j])) {
		return -1, errors.New("not a digit")
	}
	// check if we've visited this field before
	if visited[i][j] {
		return -1, errors.New("already visited")
	}

	// move left to the first digit of this number, stop when the left char of current pos is not a digit anymore
	for j-1 >= 0 && unicode.IsDigit(rune(lines[i][j-1])) {
		j--
	}

	// now move right and collect the whole number
	sb := strings.Builder{}
	for j < len(lines[i]) && unicode.IsDigit(rune(lines[i][j])) {
		visited[i][j] = true
		sb.WriteString(string(lines[i][j]))
		j++
	}

	val, err := strconv.Atoi(sb.String())
	if err != nil {
		return -1, err
	}

	return val, nil
}
