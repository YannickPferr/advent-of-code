package main

import (
	"fmt"
	"github.com/YannickPferr/advent-of-code/2023/utils"
)

func main() {
	lines, err := utils.ReadFile("2023/day3/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return
	}

	fmt.Printf("Part 1: %d\n", AnswerPart1(lines))
	fmt.Printf("Part 2: %d\n", AnswerPart2(lines))
}
