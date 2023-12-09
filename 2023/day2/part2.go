package main

func AnswerPart2(lines []string) int {
	sum := 0
	for _, line := range lines {
		_, maxRed, maxGreen, maxBlue := parseGameInfo(line)
		sum += maxRed * maxGreen * maxBlue
	}

	return sum
}
