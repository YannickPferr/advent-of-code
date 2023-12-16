package main

func AnswerPart2_2(lines []string) int {
	count := map[int]int{}
	for idx, _ := range lines {
		count[idx+1] = 1
	}

	sum := 0
	for _, line := range lines {
		cardNum, numMatches := getCardNumAndNumMatches(line)
		for i := 1; i <= numMatches; i++ {
			count[cardNum+i] += count[cardNum]
		}
		sum += count[cardNum]
	}

	return sum
}
