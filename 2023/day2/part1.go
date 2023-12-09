package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AnswerPart1(red, green, blue int, lines []string) int {
	sum := 0
	for _, line := range lines {
		gameNum, maxRed, maxGreen, maxBlue := parseGameInfo(line)
		if maxRed > red || maxGreen > green || maxBlue > blue {
			continue
		}
		sum += gameNum
	}

	return sum
}

func parseGameInfo(line string) (int, int, int, int) {
	line = strings.TrimLeft(line, "Game ")
	gameNum, err := strconv.Atoi(line[:strings.Index(line, ":")])
	if err != nil {
		fmt.Println("shouldn't happen")
		return 0, 0, 0, 0
	}
	line = line[strings.Index(line, ":")+2:]

	maxRed := 0
	maxGreen := 0
	maxBlue := 0
	rounds := strings.Split(line, "; ")
	for _, round := range rounds {
		draws := strings.Split(round, ", ")
		for _, draw := range draws {
			maxRed = max(maxRed, getNumFromDraw(draw, "red"))
			maxGreen = max(maxGreen, getNumFromDraw(draw, "green"))
			maxBlue = max(maxBlue, getNumFromDraw(draw, "blue"))
		}
	}

	return gameNum, maxRed, maxGreen, maxBlue
}

func getNumFromDraw(draw string, colorStr string) int {
	if !strings.Contains(draw, colorStr) {
		return 0
	}

	num, err := strconv.Atoi(strings.TrimRight(draw, fmt.Sprintf(" %s", colorStr)))
	if err != nil {
		fmt.Println("shouldn't happen")
		return 0
	}
	return num
}
