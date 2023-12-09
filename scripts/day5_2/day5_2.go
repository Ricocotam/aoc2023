package main

import (
	"fmt"
	"os"
    "math"
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

func parseData(filecontent string) ([]Range, Almanac) {
	var almanac Almanac
	lines := strings.Split(filecontent, "\n")
	seedsRe := regexp.MustCompile(`seeds: ((\d+\s*)+)`)
	nbsString := seedsRe.FindStringSubmatch(lines[0])[1]
	nbStringSlice := strings.Split(nbsString, " ")
	seeds := atoiSlice(nbStringSlice)
    var seedRanges []Range
    for i := 0; i < len(seeds); i+=2 {
        sr := Range{seeds[i], seeds[i] + seeds[i+1] - 1}
        seedRanges = append(seedRanges, sr)
    }

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

	return seedRanges, almanac
}

type Range struct {
    Start, End int
}

type ByStart []Range

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].Start < a[j].Start }

func mergeRange(ranges []Range) (mergedRanges []Range) {
	sort.Sort(ByStart(ranges))
    //fmt.Println("Before Merge", ranges)
    for i:=0; i < len(ranges);  {
        r := ranges[i]
        newRange := r
        j := i+1
        for ; j < len(ranges); j++ {
            if ranges[j].Start <= newRange.End {
                newRange.End = int(math.Max(float64(ranges[j].End), float64(newRange.End)))
            } else {
                break
            }
        }
        mergedRanges = append(mergedRanges, newRange)
        i = j
    }
    return mergedRanges
}

func nextRanges(ranges []Range, sortedNodes []Node) []Range {
    sort.Sort(ByStart(ranges))
    var retRanges []Range
    for _, range_ := range ranges {
        newRanges := findNextRange(range_, sortedNodes) 
        retRanges = append(retRanges, newRanges...)
    }
    retRanges = mergeRange(retRanges)
    fmt.Println("End of run :", retRanges)
    return retRanges
}

func findNextRange(currentRange Range, sortedNodes []Node) (nextRange []Range) {
    current := currentRange.Start
    for _, node := range sortedNodes {
        if current >= currentRange.End {
            break
        }
        
        // Reach node.Source
        if current < node.Source {
            newRange := Range{current, node.Source - 1}
            nextRange = append(nextRange, newRange)
            current = node.Source
            // Dont break cause now we can use current node
        }

        // Now we know current \in [node.Source, node.Source + node.Range - 1]
        maxIncluded := node.Source + node.Range - 1

        // Reach node.Source + node.Range - 1
        if current <= maxIncluded {
            newStart := current + node.Dest - node.Source
            newEnd := maxIncluded + node.Dest - node.Source
            maxEnd := currentRange.End + node.Dest - node.Source
            if maxEnd < newEnd {
                newEnd = maxEnd
            }
            newRange := Range{newStart, newEnd}
            nextRange = append(nextRange, newRange)
            current = current + newEnd - newStart + 1
        }
    }

    // complete untouched 
    if current <= currentRange.End { 
        newRange := Range{current, currentRange.End}
        nextRange = append(nextRange, newRange)
        current = current + newRange.End - newRange.Start + 1
        //fmt.Println("Next", nextRange, current)
    }
    return nextRange
}


func locationNumber(seed []Range, almanac Almanac) []Range {
	soil := nextRanges(seed, almanac.SeedSoil)
	fertilizer := nextRanges(soil, almanac.SoilFertilizer)
	water := nextRanges(fertilizer, almanac.FertilizerWater)
	light := nextRanges(water, almanac.WaterLight)
	temperature := nextRanges(light, almanac.LightTemperature)
	humidity := nextRanges(temperature, almanac.TemperatureHumidity)
	location := nextRanges(humidity, almanac.HumidityLocation)

    //fmt.Println(seed, soil, fertilizer, water, light, temperature, humidity, location)

	return location
}

func main() {
	filecontent, err := os.ReadFile("data/day5_input.txt")
	if err != nil {
		panic(err)
	}
	seedRanges, almanac := parseData(string(filecontent))
    fmt.Println("Seeds :", seedRanges) 
    locationRanges := locationNumber(seedRanges, almanac)
    sort.Sort(ByStart(locationRanges))
    lowest := locationRanges[0].Start

	fmt.Println(lowest)
}
