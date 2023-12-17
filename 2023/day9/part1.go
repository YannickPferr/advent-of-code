package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AnswerPart1(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += predictNextNum(line)
	}
	return sum
}

func predictNextNum(line string) int {
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
	return completeSequence(mat)
}

func completeSequence(mat [][]int) int {
	mat[len(mat)-1] = append(mat[len(mat)-1], 0)
	for i := len(mat) - 2; i >= 0; i-- {
		val := mat[i][len(mat[i])-1] + mat[i+1][len(mat[i+1])-1]
		mat[i] = append(mat[i], val)
	}
	return mat[0][len(mat[0])-1]
}

func buildSequence(mat [][]int, idx int) [][]int {
	var row []int
	input := mat[idx]
	allZeroes := true
	for i := 1; i < len(input); i++ {
		val := input[i] - input[i-1]
		if val != 0 {
			allZeroes = false
		}
		row = append(row, val)
	}
	mat = append(mat, row)
	if allZeroes {
		return mat
	}

	return buildSequence(mat, idx+1)
}
