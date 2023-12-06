package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CardPile struct {
	Winning, Mine map[int]bool
}

func buildPile(winningPile string, myPile string) (cardPile CardPile) {
	cardPile.Winning = make(map[int]bool)
	cardPile.Mine = make(map[int]bool)

	for _, nb := range strings.Split(winningPile, " ") {
		nbInt, err := strconv.Atoi(nb)
		if err != nil {
			continue
		}
		cardPile.Winning[nbInt] = true
	}

	for _, nb := range strings.Split(myPile, " ") {
		nbInt, err := strconv.Atoi(nb)
		if err != nil {
			continue
		}
		cardPile.Mine[nbInt] = true
	}

	return cardPile
}

func parse_data(filename string) (cardPiles []CardPile) {
	fileContent, err := os.ReadFile(filename)
	fc := string(fileContent)

	if err != nil {
		panic(err)
	}

	dataRe := regexp.MustCompile(`(?P<winning>(\s+\d+)*) \|(?P<mine>(\s+\d+)*)`)
	lines := strings.Split(fc, "\n")

	winning_idx := dataRe.SubexpIndex("winning")
	mine_idx := dataRe.SubexpIndex("mine")
	for _, line := range lines {
		res := dataRe.FindStringSubmatch(line)
		winning := res[winning_idx]
		mine := res[mine_idx]

		pile := buildPile(winning, mine)
		cardPiles = append(cardPiles, pile)
	}
	return cardPiles
}

func nbIntersection(cp CardPile) int {
	var result int = 0
	for key := range cp.Winning {
		if cp.Mine[key] {
			result += 1
		}
	}
	return result
}

func main() {
	cardPiles := parse_data("data/day4_input.txt")

	var total int = 0
	var nbTimes []int = make([]int, len(cardPiles))
	for i := 0; i < len(nbTimes); i++ {
		nbTimes[i] = 1
	}

	for i, cardPile := range cardPiles {
		nb := nbIntersection(cardPile)
		for j := i + 1; j < i+nb+1; j++ {
			nbTimes[j] += nbTimes[i]
		}
		total += nbTimes[i]
	}
	fmt.Println(total)
}
