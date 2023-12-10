package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AnswerPart1(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += getGameScore(line)
	}

	return sum
}

func getGameScore(line string) int {
	line = line[strings.Index(line, ": ")+2:]
	lists := strings.Split(line, " | ")

	winningCards := map[int]bool{}
	for _, winningCard := range strings.Fields(lists[0]) {
		val, err := strconv.Atoi(winningCard)
		if err != nil {
			fmt.Println("shouldn't happen")
		}
		winningCards[val] = true
	}

	score := 0
	for _, ownedCard := range strings.Fields(lists[1]) {
		val, err := strconv.Atoi(ownedCard)
		if err != nil {
			fmt.Println("shouldn't happen")
		}

		if winningCards[val] {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	return score
}
