package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func compute_line(line string) int {
	var res int = 0

	for i := 0; i < len(line); i++ {
		tmp, err := strconv.Atoi(string(line[i]))
		if err != nil {
			continue
		}
		res += 10 * tmp
		break
	}
	for i := 0; i < len(line); i++ {
		j := len(line) - i - 1
		tmp, err := strconv.Atoi(string(line[j]))
		if err != nil {
			continue
		}
		res += tmp
		break
	}

	return res
}

var stringNumber map[string]string = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func replacer(number string) string {
	return stringNumber[number]
}

func transformStringNumbers(text string) string {

	myregex := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine")

	finalText := myregex.ReplaceAllStringFunc(text, replacer)
	return finalText
}

func main() {
	raw_data, err := os.ReadFile("data/test_input.txt")
	if err != nil {
		panic(err)
	}

	data := string(raw_data)
	data = transformStringNumbers(data)
	fmt.Println(data)
	lines := strings.Split(data, "\n")

	var values []int = make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		res := compute_line(lines[i])
		values[i] = res
	}

	var total int = 0
	for i := 0; i < len(lines); i++ {
		total += values[i]
	}

	fmt.Println(total)

}
