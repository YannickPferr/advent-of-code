package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AnswerPart2(lines []string) int {
	game := parseGame(lines)
	lowerBound := binarySearchLowerBound(game)
	upperBound := binarySearchUpperBound(game)
	return upperBound + 1 - lowerBound
}

func parseGame(lines []string) Game {
	time := strings.Join(strings.Fields(strings.TrimLeft(lines[0], "Time:")), "")
	record := strings.Join(strings.Fields(strings.TrimLeft(lines[1], "Distance:")), "")
	timeInt, err := strconv.Atoi(time)
	if err != nil {
		fmt.Println("shouldn't happen")
	}
	recordInt, err := strconv.Atoi(record)
	if err != nil {
		fmt.Println("shouldn't happen")
	}
	return Game{
		time:   timeInt,
		record: recordInt,
	}
}
