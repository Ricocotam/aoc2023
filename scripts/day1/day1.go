package main

import (
	"fmt"
	"os"
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
	fmt.Println(res)
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

func main() {
	raw_data, err := os.ReadFile("data/day1_input.txt")
	if err != nil {
		panic(err)
	}

	data := string(raw_data)
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
