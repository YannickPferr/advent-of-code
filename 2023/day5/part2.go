package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func AnswerPart2(lines []string) int {
	seedRanges := getSeedsWithRanges(lines)
	maps := getMaps(lines)

	minLocation := math.MaxInt64
	for _, seed := range seedRanges {
		convertedVal := []Range{seed}
		for _, m := range maps {
			convertedVal = m.ConvertRange(convertedVal)
		}

		minLocation = min(minLocation, convertSeedRange(seed, maps))
	}
	return minLocation
}

func getSeedsWithRanges(lines []string) []Range {
	seedsLine := strings.TrimLeft(lines[0], "seeds: ")
	seedsStr := strings.Fields(seedsLine)

	var seedRanges []Range
	for i := 0; i+1 < len(seedsStr); i += 2 {
		seedRangeStart, err := strconv.Atoi(seedsStr[i])
		if err != nil {
			fmt.Println("shouldn't happen")
			continue
		}
		seedRangeLength, err := strconv.Atoi(seedsStr[i+1])
		if err != nil {
			fmt.Println("shouldn't happen")
			continue
		}

		seedRanges = append(seedRanges, NewSeedRange(seedRangeStart, seedRangeLength))
	}
	return seedRanges
}

func convertSeedRange(seed Range, maps []Map) int {
	convertedVal := []Range{seed}
	for _, m := range maps {
		convertedVal = m.ConvertRange(convertedVal)
	}

	minStart := math.MaxInt
	for _, r := range convertedVal {
		minStart = min(minStart, r.start)
	}
	return minStart
}
