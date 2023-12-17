package main

import (
	"strings"
)

func AnswerPart2(lines []string) int {
	instructions := lines[0]

	graph := map[string][]string{}
	var startingNodes []string
	for i := 2; i < len(lines); i++ {
		node, neighbors := getNodeAndNeighbors(lines[i])
		graph[node] = neighbors

		if strings.HasSuffix(node, "A") {
			startingNodes = append(startingNodes, node)
		}
	}

	var numStepsArr []int
	for _, startingNode := range startingNodes {
		numSteps := 0
		for !strings.HasSuffix(startingNode, "Z") {
			for _, instruction := range instructions {
				neighbor := 0
				if instruction == 'R' {
					neighbor = 1
				}
				startingNode = graph[startingNode][neighbor]
				numSteps++
			}
		}
		numStepsArr = append(numStepsArr, numSteps)
	}

	return getLeastCommonDenominator(numStepsArr)
}

func getLeastCommonDenominator(denominators []int) int {
	// Find the prime factors of each denominator
	primeFactors := make(map[int]int)
	for _, d := range denominators {
		factors := primeFactorization(d)
		for factor, count := range factors {
			if primeFactors[factor] < count {
				primeFactors[factor] = count
			}
		}
	}

	// Multiply the prime factors together to get the LCD
	lcd := 1
	for factor, count := range primeFactors {
		for i := 0; i < count; i++ {
			lcd *= factor
		}
	}
	return lcd
}

// Returns a map of prime factors and their counts for a given number
func primeFactorization(n int) map[int]int {
	factors := make(map[int]int)
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors[i]++
			n /= i
		}
	}
	return factors
}
