package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card int64

const (
	Two Card = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	As
)

type HandForce int32

const (
	HighCard HandForce = iota
	OnePair
	TwoPairs
	Brelan
	Full
	Carre
	SuperCarre
)

type Hand struct {
	Cards [5]Card
	Bid   int
}

type ByCards []Hand

func (a ByCards) Len() int      { return len(a) }
func (a ByCards) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func getHandForce(cards [5]Card) (force HandForce) {
	var counter [13]int
	force = -1
	for _, c := range cards {
		counter[c] += 1
	}
	hasThree, hasPair := false, false
	for _, c := range counter {
		if c == 2 {
			if !hasPair {
				hasPair = true
			} else {
				force = TwoPairs
				break
			}
		} else if c == 3 {
			hasThree = true
		} else if c == 4 {
			force = Carre
			break
		} else if c == 5 {
			force = SuperCarre
			break
		}
	}
	if force == -1 {
		if hasThree {
			if hasPair {
				force = Full
			} else {
				force = Brelan
			}
		} else if hasPair {
            force = OnePair
        } else {
            force = HighCard
        }
	}
	return force
}

func (a ByCards) Less(i, j int) bool {
	forceI := getHandForce(a[i].Cards)
	forceJ := getHandForce(a[j].Cards)
	if forceI < forceJ {
		return true
	} else if forceJ < forceI {
		return false
	}

	for cI := 0; cI < 5; cI++ {
		if a[i].Cards[cI] < a[j].Cards[cI] {
			return true
		} else if a[i].Cards[cI] > a[j].Cards[cI] {
			return false
		}
	}
	panic("Error")
}

func parseHand(line string) (hand Hand) {
	handBid := strings.Split(line, " ")
	bidString := handBid[1]
	bid, err := strconv.Atoi(bidString)
	if err != nil {
		panic(err)
	}
	hand.Bid = bid

	for i, c := range handBid[0] {
		card := Card(-1)
		switch c {
		case '2':
			card = Two
		case '3':
			card = Three
		case '4':
			card = Four
		case '5':
			card = Five
		case '6':
			card = Six
		case '7':
			card = Seven
		case '8':
			card = Eight
		case '9':
			card = Nine
		case 'T':
			card = Ten
		case 'J':
			card = Jack
		case 'Q':
			card = Queen
		case 'K':
			card = King
		case 'A':
			card = As
		default:
			panic("Error card doesn't exist")
		}
		hand.Cards[i] = card
	}

	return hand
}

func parseData(filename string) (hands []Hand) {
	filecontent, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(filecontent), "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		hand := parseHand(line)
		hands = append(hands, hand)
	}

	return hands
}

func main() {
    var total int = 0
	hands := parseData("data/day7_input.txt")

	sort.Sort(ByCards(hands))
    for i, hand := range hands {
        total += (i+1) * hand.Bid
    }
    fmt.Println(total)
}
