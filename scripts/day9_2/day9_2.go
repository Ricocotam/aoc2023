package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Seq []int

func atoiSlice(ss []string) (is Seq) {
	for _, s := range ss {
		s = strings.TrimSpace(s)
		i, err := strconv.Atoi(s)
		if err != nil {

			panic(err)
		}
		is = append(is, i)
	}
	return is
}

func parseData(filename string) []Seq {
	filecontent, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(filecontent), "\n")
	var sequences []Seq
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		splitted := strings.Split(line, " ")
        sequences = append(sequences, atoiSlice(splitted))
	}
	return sequences
}

func nullSeq(seq Seq) bool {
	for _, e := range seq {
		if e != 0 {
			return false
		}
	}
	return true
}

func nextSeq(seq Seq) Seq {
	var newSeq []int = make([]int, len(seq)-1)
	for i := 0; i < len(seq)-1; i++ {
		newSeq[i] = seq[i+1] - seq[i]
	}
	return Seq(newSeq)
}

func computeValue(firsts []int) (value int) {
	value = 0
    fmt.Print(firsts, ";")
    for i := len(firsts)-2; i >= 0; i-- {
        first := firsts[i]
        fmt.Print(first, value, ",")
		value = first - value
	}
    fmt.Println()
	return value
}

func predictValue(seq Seq) (value int) {
	var firsts []int
    firsts = append(firsts, seq[0])
	for !nullSeq(seq) {
		seq = nextSeq(seq)
		firsts = append(firsts, seq[0])
	}
	value = computeValue(firsts)
	return value
}

func main() {
	data := parseData("data/day9_input.txt")
	var total int = 0
	for _, d := range data {
		value := predictValue(d)
		total += value
	}
	fmt.Println(total)
}
