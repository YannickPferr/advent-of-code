package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AnswerPart2(lines []string) int {
	// build graph
	graph := map[int][]int{}
	for _, line := range lines {
		cardNum, numMatches := getCardNumAndNumMatches(line)
		var connections []int
		for i := 1; i <= numMatches; i++ {
			connections = append(connections, cardNum+i)
		}
		graph[cardNum] = connections
	}

	// init cache with -1
	cache := make([]int, len(graph)+1)
	for idx := range cache {
		cache[idx] = -1
	}
	// dfs from every node in graph and count num nodes visited (cache visited nodes)
	nodesVisited := 0
	for currentNode, _ := range graph {
		nodesVisited += dfs(graph, currentNode, cache)
	}

	return nodesVisited
}

func getCardNumAndNumMatches(line string) (int, int) {
	line = strings.TrimLeft(line, "Card")
	line = strings.TrimLeft(line, " ")
	cardNum, err := strconv.Atoi(line[:strings.Index(line, ":")])
	if err != nil {
		fmt.Println("shouldn't happen")
	}

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

	numMatches := 0
	for _, ownedCard := range strings.Fields(lists[1]) {
		val, err := strconv.Atoi(ownedCard)
		if err != nil {
			fmt.Println("shouldn't happen")
		}

		if winningCards[val] {
			numMatches++
		}
	}
	return cardNum, numMatches
}

func dfs(graph map[int][]int, currentNode int, cache []int) int {
	if cache[currentNode] != -1 {
		return cache[currentNode]
	}

	sumVisits := 1
	for _, neighbor := range graph[currentNode] {
		sumVisits += dfs(graph, neighbor, cache)
	}
	cache[currentNode] = sumVisits
	return sumVisits
}
