package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readData(filename string) []string {
	raw_data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := string(raw_data)
	lines := strings.Split(data, "\n")
	return lines
}

type gameData struct {
	Red, Green, Blue int
}

func (gd gameData) allLowerOrEqualThan(other gameData) bool {
	return gd.Red <= other.Red && gd.Green <= other.Green && gd.Blue <= other.Blue
}

var gameIdRe regexp.Regexp = *regexp.MustCompile(`Game\s(?P<gameId>\d+)`)
var redNumbersRe regexp.Regexp = *regexp.MustCompile(`(?P<number>\d+) red`)
var greenNumbersRe regexp.Regexp = *regexp.MustCompile(`(?P<number>\d+) green`)
var blueNumbersRe regexp.Regexp = *regexp.MustCompile(`(?P<number>\d+) blue`)

func maxIntSlice(intSlice []int) int {
	var max int = 0

	for i := 0; i < len(intSlice); i++ {
		if intSlice[i] > max {
			max = intSlice[i]
		}
	}

	return max
}

func atoiSlice(aSlice [][]string) []int {
	res := make([]int, len(aSlice))
	var err error
	for i := 0; i < len(aSlice); i++ {
		// aSlice[i][1]
		// 1 because SubmatchIndex["number"] = 1
		res[i], err = strconv.Atoi(string(aSlice[i][1]))
		if err != nil {
			panic(err)
		}
	}
	return res
}

func buildGameData(line string) (gameData, int) {
	matches := gameIdRe.FindStringSubmatch(line)
	gameIdIndex := gameIdRe.SubexpIndex("gameId")
	gameId, err := strconv.Atoi(matches[gameIdIndex])

	if err != nil {
		panic(err)
	}

	redStringSlice := redNumbersRe.FindAllStringSubmatch(line, -1)
	greenStringSlice := greenNumbersRe.FindAllStringSubmatch(line, -1)
	blueStringSlice := blueNumbersRe.FindAllStringSubmatch(line, -1)

	redIntSlice := atoiSlice(redStringSlice)
	greenIntSlice := atoiSlice(greenStringSlice)
	blueIntSlice := atoiSlice(blueStringSlice)

	red := maxIntSlice(redIntSlice)
	green := maxIntSlice(greenIntSlice)
	blue := maxIntSlice(blueIntSlice)

	return gameData{red, green, blue}, gameId
}

func main() {
	var maxGameData gameData = gameData{12, 13, 14}
	lines := readData("data/day2_input.txt")
	var total int = 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		gd, gId := buildGameData(line)
		if gd.allLowerOrEqualThan(maxGameData) {
			total += gId
		}
	}
	fmt.Println(total)

}
