package main

import (
	"sort"
)

func AnswerPart2(lines []string) int {
	var hands []Hand
	for _, line := range lines {
		hands = append(hands, parseHand(line, true))
	}

	// sort
	sort.Slice(hands, func(i, j int) bool {
		handA := hands[i]
		handB := hands[j]

		if handA.handType == handB.handType {
			for i := 0; i < len(handA.cards); i++ {
				symbolA := handA.cards[i]
				symbolB := handB.cards[i]
				if symbolA == symbolB {
					continue
				}
				return symbolA < symbolB
			}
		}

		return handA.handType < handB.handType
	})

	sum := 0
	for idx, hand := range hands {
		sum += hand.bidAmount * (idx + 1)
	}
	return sum
}
