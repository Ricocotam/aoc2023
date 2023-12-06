package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	Dest, Source, Range int
}

type BySource []Node

func (a BySource) Len() int           { return len(a) }
func (a BySource) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySource) Less(i, j int) bool { return a[i].Source < a[j].Source }

type Almanac struct {
	SeedSoil            []Node
	SoilFertilizer      []Node
	FertilizerWater     []Node
	WaterLight          []Node
	LightTemperature    []Node
	TemperatureHumidity []Node
	HumidityLocation    []Node
}

func atoiSlice(ss []string) (is []int) {
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		is = append(is, i)
	}
	return is
}

func parseMapping(lines []string, start int) ([]Node, int) {
	var mappingData []Node
	for ; start < len(lines) && len(lines[start]) > 0; start++ {
		nbStringSlice := strings.Split(lines[start], " ")
		nbs := atoiSlice(nbStringSlice)

		node := Node{nbs[0], nbs[1], nbs[2]}
		mappingData = append(mappingData, node)
	}
	sort.Sort(BySource(mappingData))
	return mappingData, start
}

func parseData(filecontent string) ([]int, Almanac) {
	var almanac Almanac
	lines := strings.Split(filecontent, "\n")
	seedsRe := regexp.MustCompile(`seeds: ((\d+\s*)+)`)
	nbsString := seedsRe.FindStringSubmatch(lines[0])[1]
	nbStringSlice := strings.Split(nbsString, " ")
	seeds := atoiSlice(nbStringSlice)

	i := 3
	almanac.SeedSoil, i = parseMapping(lines, i)

	i += 2
	almanac.SoilFertilizer, i = parseMapping(lines, i)

	i += 2
	almanac.FertilizerWater, i = parseMapping(lines, i)

	i += 2
	almanac.WaterLight, i = parseMapping(lines, i)

	i += 2
	almanac.LightTemperature, i = parseMapping(lines, i)

	i += 2
	almanac.TemperatureHumidity, i = parseMapping(lines, i)

	i += 2
	almanac.HumidityLocation, i = parseMapping(lines, i)

	return seeds, almanac
}

func findNext(current int, sortedNodes []Node) (next int) {
	next = -1
	for i, node := range sortedNodes {
		if current < node.Source {
			next = current
			break
		}
		max := node.Source + node.Range
		// < or <= if max - 1
		if current < max {
			next = node.Dest + current - node.Source
			break
		} else if i+1 < len(sortedNodes) && current < sortedNodes[i+1].Source {
			next = current
			break
		}
	}

	return next
}

func locationNumber(seed int, almanac Almanac) int {
	soil := findNext(seed, almanac.SeedSoil)
	fertilizer := findNext(soil, almanac.SoilFertilizer)
	water := findNext(fertilizer, almanac.FertilizerWater)
	light := findNext(water, almanac.WaterLight)
	temperature := findNext(light, almanac.LightTemperature)
	humidity := findNext(temperature, almanac.TemperatureHumidity)
	location := findNext(humidity, almanac.HumidityLocation)

	fmt.Println(seed, soil, fertilizer, water, light, temperature, humidity, location)

	return location
}

func main() {
	filecontent, err := os.ReadFile("data/day5_input.txt")
	if err != nil {
		panic(err)
	}
	seeds, almanac := parseData(string(filecontent))
	loc := locationNumber(seeds[0], almanac)
	lowest := loc
	for _, seed := range seeds[1:] {
		loc = locationNumber(seed, almanac)
		if loc < lowest {
			lowest = loc
		}
	}
	fmt.Println(lowest)
}
