package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func AnswerPart1(lines []string) int {
	seeds := getSeeds(lines)
	maps := getMaps(lines)

	minLocation := math.MaxInt64
	for _, seed := range seeds {
		minLocation = min(minLocation, convert(seed, maps))
	}
	return minLocation
}

func getSeeds(lines []string) []int {
	seedsLine := strings.TrimLeft(lines[0], "seeds: ")
	seedsStr := strings.Fields(seedsLine)

	var seeds []int
	for _, seedStr := range seedsStr {
		intVal, err := strconv.Atoi(seedStr)
		if err != nil {
			fmt.Println("shouldn't happen")
			continue
		}
		seeds = append(seeds, intVal)
	}
	return seeds
}

func getMaps(lines []string) []Map {
	var maps []Map

	mapLines := lines[2:]
	currentMap := Map{}
	for _, line := range mapLines {
		if strings.Contains(line, ":") {
			continue
		}

		if line == "" {
			maps = append(maps, currentMap)
			currentMap = Map{}
			continue
		}

		currentMap.AddRule(getConversionRule(line))
	}
	maps = append(maps, currentMap)
	currentMap = Map{}
	return maps
}

func getConversionRule(line string) ConversionRule {
	fields := strings.Fields(line)

	dstStart, err := strconv.Atoi(fields[0])
	if err != nil {
		fmt.Println("shouldn't happen")
	}

	srcStart, err := strconv.Atoi(fields[1])
	if err != nil {
		fmt.Println("shouldn't happen")
	}

	rangeLength, err := strconv.Atoi(fields[2])
	if err != nil {
		fmt.Println("shouldn't happen")
	}

	return NewConversionRule(srcStart, dstStart, rangeLength)
}

func convert(seed int, maps []Map) int {
	convertedVal := seed
	for _, m := range maps {
		convertedVal = m.Convert(convertedVal)
	}
	return convertedVal
}

//seed-to-soil map:
//soil-to-fertilizer map:
//fertilizer-to-water map:
//water-to-light map:
//light-to-temperature map:
//temperature-to-humidity map:
//humidity-to-location map:
