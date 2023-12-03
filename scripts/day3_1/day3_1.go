package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}

type PuzzleData struct {
	NumbersMap map[Coordinate]int
	PuzzleMap  [][]byte
}

func parse_data(filename string) (data PuzzleData) {
	// Read Data
	byteContent, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	stringContent := string(byteContent)
	lines := strings.Split(stringContent, "\n")

	data.PuzzleMap = make([][]byte, len(lines))
	for i, line := range lines {
		data.PuzzleMap[i] = []byte(line)
	}

	// Build NumbersMap
	// Find numbers and build a map that maps for each digit's coordinate
	// associated number
	data.NumbersMap = make(map[Coordinate]int)
	digitRe := regexp.MustCompile(`\d+`)
	for x, line := range data.PuzzleMap {
		indexes := digitRe.FindAllIndex(line, -1)
		for _, idxs := range indexes {
			nbString := string(line[idxs[0]:idxs[len(idxs)-1]])
			nb, err := strconv.Atoi(string(nbString))
			if err != nil {
				continue
			}
			for y := idxs[0]; y < idxs[len(idxs)-1]; y++ {
				coord := Coordinate{x, y}
				data.NumbersMap[coord] = nb
			}
		}
	}

	return data
}

func get_numbers(data PuzzleData) (numbers []int) {
	symbolRe := regexp.MustCompile(`[^\d\.]`)
	for x, line := range data.PuzzleMap {
		indexes := symbolRe.FindAllIndex(line, -1)
		for _, idxs := range indexes {
			y := idxs[0]
			var prevNb int = -1
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					coord := Coordinate{x + i, y + j}
					if nb, ok := data.NumbersMap[coord]; ok {
						if nb != prevNb {
							numbers = append(numbers, nb)
						}
						prevNb = nb
					} else {
						prevNb = -1
					}
				}
			}
		}
	}
	return numbers
}

func main() {
	var total int = 0

	data := parse_data("data/day3_input.txt")
	fmt.Println(data)
	numbers := get_numbers(data)

	for _, nb := range numbers {
		total += nb
	}

	fmt.Println(total)
}
