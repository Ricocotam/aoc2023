package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func atoiSlice(ss []string) (is []int) {
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
            continue
		}
		is = append(is, i)
	}
	return is
}

func parseData(filename string) (times []int, records []int) {
    filecontent, err := os.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    timesRe := regexp.MustCompile(`Time:\s+(?P<times>(\d+\s+)+)`)
    distanceRe := regexp.MustCompile(`Distance:\s+(?P<distance>(\d+\s+)+)`)
    
    timesString := timesRe.FindStringSubmatch(string(filecontent))[1]
    timesString = strings.ReplaceAll(timesString, "\n", " ")
    timesString = strings.ReplaceAll(timesString, " ", "")
    timesStringSlice := strings.Split(timesString, " ")
    times = atoiSlice(timesStringSlice)

    distanceString := distanceRe.FindStringSubmatch(string(filecontent))[1]
    distanceString = strings.ReplaceAll(distanceString, "\n", " ")
    distanceString = strings.ReplaceAll(distanceString, " ", "")
    distanceStringSlice := strings.Split(distanceString, " ")
    records = atoiSlice(distanceStringSlice)

    return times, records
}

func compute(time int, record int) (nb int) {
    sqrtDelta := math.Sqrt(math.Pow(float64(time), 2) - 4 * float64(record))
    x1 := (time - int(sqrtDelta)) / 2 
    x2 := (time + int(sqrtDelta)) / 2 
    nb = x2 - x1 + 1

    if x1 * (time - x1) <= record {
        nb -= 1
    }
    if x2 * (time - x2) <= record {
        nb -= 1
    }

    fmt.Println(x1, x2, nb)
    return nb
}

func main() {
    var total int = 1
    times, records := parseData("data/day6_input.txt")
    for i:=0;i < len(times); i++ {
        time := times[i]
        record := records[i]
        nb := compute(time, record)
        total *= nb
    }
    
    fmt.Println(total)
}
