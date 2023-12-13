package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	time   int
	record int
}

func AnswerPart1(lines []string) int {
	games := parseGames(lines)
	answer := 1
	for _, game := range games {
		lowerBound := binarySearchLowerBound(game)
		upperBound := binarySearchUpperBound(game)
		answer *= upperBound + 1 - lowerBound
	}
	return answer
}

func binarySearchLowerBound(game Game) int {
	left := 0
	right := game.time
	for left < right {
		mid := left + (right-left)/2
		if isAbleToBeatRecord(game.record, game.time, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func isAbleToBeatRecord(record int, gameTime int, holdButtonTime int) bool {
	return holdButtonTime*(gameTime-holdButtonTime) > record
}

func binarySearchUpperBound(game Game) int {
	left := 0
	right := game.time
	for left < right {
		mid := left + (right-left)/2
		if isAbleToBeatRecord(game.record, game.time, mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left - 1
}

func parseGames(lines []string) []Game {
	times := strings.Fields(strings.TrimLeft(lines[0], "Time:"))
	records := strings.Fields(strings.TrimLeft(lines[1], "Distance:"))
	var games []Game
	for idx, time := range times {
		timeInt, err := strconv.Atoi(time)
		if err != nil {
			fmt.Println("shouldn't happen")
			continue
		}
		recordInt, err := strconv.Atoi(records[idx])
		if err != nil {
			fmt.Println("shouldn't happen")
			continue
		}
		games = append(games, Game{
			time:   timeInt,
			record: recordInt,
		})
	}
	return games
}
