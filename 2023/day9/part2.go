package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AnswerPart2(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += predictPreviousNum(line)
	}
	return sum
}

func predictPreviousNum(line string) int {
	var mat [][]int
	var row []int
	for _, num := range strings.Fields(line) {
		int, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("shouldn't happen")
		}
		row = append(row, int)
	}
	mat = append(mat, row)
	mat = buildSequence(mat, 0)
	return completePreviousSequence(mat)
}

func completePreviousSequence(mat [][]int) int {
	mat[len(mat)-1] = append([]int{0}, mat[len(mat)-1]...)
	for i := len(mat) - 2; i >= 0; i-- {
		val := mat[i][0] - mat[i+1][0]
		mat[i] = append([]int{val}, mat[i]...)
	}
	return mat[0][0]
}
