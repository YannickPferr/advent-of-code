package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	handType  HandType
	cards     []int
	bidAmount int
}

func AnswerPart1(lines []string) int {
	var hands []Hand
	for _, line := range lines {
		hands = append(hands, parseHand(line, false))
	}

	// sort
	sort.Slice(hands, func(i, j int) bool {
		handA := hands[i]
		handB := hands[j]

		if handA.handType == handB.handType {
			// compare individual cards
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

func parseHand(line string, withJoker bool) Hand {
	hand := strings.Fields(line)
	bidAmount, err := strconv.Atoi(hand[1])
	if err != nil {
		fmt.Println("shouldn't happen")
	}
	return Hand{
		handType:  getHandType(hand[0], withJoker),
		cards:     transformHandToIntArray(hand[0], withJoker),
		bidAmount: bidAmount,
	}
}

func getHandType(cards string, withJoker bool) HandType {
	cardMap := map[rune]int{}
	jokerCount := 0
	for _, cardSymbol := range cards {
		if withJoker && cardSymbol == 'J' {
			jokerCount++
			continue
		}

		_, exists := cardMap[cardSymbol]
		if !exists {
			cardMap[cardSymbol] = 1
			continue
		}
		cardMap[cardSymbol]++
	}

	// all cards are the same = FiveOfAKind
	if len(cardMap) <= 1 {
		return FiveOfAKind
	}

	// 2 different cards = FourOfAKind or FullHouse
	if len(cardMap) == 2 {
		// if one of the counts is a 4 = FourOfAKind
		for _, count := range cardMap {
			if count+jokerCount == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	}

	// 3 different cards = ThreeOfAKind or TwoPair
	if len(cardMap) == 3 {
		// if one of the counts is a 3 = ThreeOfAKind
		for _, count := range cardMap {
			if count+jokerCount == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	}

	// 4 different cards = OnePair
	if len(cardMap) == 4 {
		return OnePair
	}

	// 5 different cards = HighCard
	return HighCard
}

func transformHandToIntArray(hand string, withJoker bool) []int {
	var ints []int
	for _, symbol := range hand {
		switch symbol {
		case '2':
			ints = append(ints, 2)
		case '3':
			ints = append(ints, 3)
		case '4':
			ints = append(ints, 4)
		case '5':
			ints = append(ints, 5)
		case '6':
			ints = append(ints, 6)
		case '7':
			ints = append(ints, 7)
		case '8':
			ints = append(ints, 8)
		case '9':
			ints = append(ints, 9)
		case 'T':
			ints = append(ints, 10)
		case 'J':
			if withJoker {
				ints = append(ints, 1)
			} else {
				ints = append(ints, 11)
			}
		case 'Q':
			ints = append(ints, 12)
		case 'K':
			ints = append(ints, 13)
		case 'A':
			ints = append(ints, 14)
		}
	}
	return ints
}
