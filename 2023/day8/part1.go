package main

import "strings"

func AnswerPart1(lines []string) int {
	instructions := lines[0]

	graph := map[string][]string{}
	for i := 2; i < len(lines); i++ {
		node, neighbors := getNodeAndNeighbors(lines[i])
		graph[node] = neighbors
	}

	numSteps := 0
	currentNode := "AAA"
	for currentNode != "ZZZ" {
		for _, instruction := range instructions {
			neighbor := 0
			if instruction == 'R' {
				neighbor = 1
			}
			currentNode = graph[currentNode][neighbor]
			numSteps++
		}
	}

	return numSteps
}

func getNodeAndNeighbors(line string) (string, []string) {
	nodeAndNeighbor := strings.Split(line, "=")
	neighborsStr := strings.TrimSpace(nodeAndNeighbor[1])
	neighborsStr = strings.TrimLeft(neighborsStr, "(")
	neighborsStr = strings.TrimRight(neighborsStr, ")")
	neighbors := strings.Split(neighborsStr, ", ")
	return strings.TrimSpace(nodeAndNeighbor[0]), neighbors
}
