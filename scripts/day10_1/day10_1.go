package main

import (
	"fmt"
	"os"
	"strings"
)

type Pipe string

const (
	NorthSouth Pipe = "|"
	EastWest   Pipe = "-"
	NorthEast  Pipe = "L"
	NorthWest  Pipe = "J"
	SouthWest  Pipe = "7"
	SouthEast  Pipe = "F"
	Start      Pipe = "S"
	Ground     Pipe = "."
)

type Graph [][]Pipe

func parsedata(filename string) (graph Graph, startingPosition [2]int) {
	filecontent, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(filecontent), "\n")
    graph = make([][]Pipe, len(lines)-1)
    for i, line := range lines {
        if len(line) == 0 {
            continue
        }
		graph[i] = make([]Pipe, len(line))
		for j, b := range line {
			graph[i][j] = Pipe(b)
			if graph[i][j] == Start {
				startingPosition = [2]int{i, j}
			}
		}
	}
	return graph, startingPosition
}

func different(pos1, pos2 [2]int) bool {
	return pos1[0] != pos2[0] || pos1[1] != pos2[1]
}

func isNorth(pos [2]int, graph Graph) bool {
    if !isValid(pos, graph) {
        return false
    }
	pipe := Pipe(graph[pos[0]][pos[1]])
	return pipe == NorthSouth || pipe == NorthWest || pipe == NorthEast
}

func isSouth(pos [2]int, graph Graph) bool {
    if !isValid(pos, graph) {
        return false
    }
	pipe := Pipe(graph[pos[0]][pos[1]])
	return pipe == NorthSouth || pipe == SouthWest || pipe == SouthEast
}

func isEast(pos [2]int, graph Graph) bool {
    if !isValid(pos, graph) {
        return false
    }
	pipe := Pipe(graph[pos[0]][pos[1]])
	return pipe == EastWest || pipe == NorthEast || pipe == SouthEast
}

func isWest(pos [2]int, graph Graph) bool {
    if !isValid(pos, graph) {
        return false
    }
	pipe := Pipe(graph[pos[0]][pos[1]])
	return pipe == EastWest || pipe == NorthWest || pipe == SouthWest
}

func isValid(pos [2]int, graph Graph) bool {
	nonNegative := pos[0] > -1 && pos[1] > -1
	notTooBig := pos[0] < len(graph) && pos[1] < len(graph[0])
	return nonNegative && notTooBig
}

func replaceStart(pos [2]int, graph Graph) Pipe {
	east, west, north, south := false, false, false, false

	graph[pos[0]][pos[1]] = NorthSouth
	if isSouth([2]int{pos[0] - 1, pos[1]}, graph) {
		north = true
	}
	if isNorth([2]int{pos[0] + 1, pos[1]}, graph) {
		south = true
	}
	graph[pos[0]][pos[1]] = EastWest
	if isWest([2]int{pos[0], pos[1] + 1}, graph) {
		east = true
	}
	if isWest([2]int{pos[0], pos[1] - 1}, graph) {
		west = true
	}
	newStart := Pipe("")
	if north && south {
		newStart = NorthSouth
	} else if north && east {
		newStart = NorthEast
	} else if north && west {
		newStart = NorthWest
	} else if south && east {
		newStart = SouthEast
	} else if south && west {
		newStart = SouthWest
	} else if east && west {
		newStart = EastWest
	}
	return newStart
}

func solve(graph Graph, startingPos [2]int) (length int) {
	length = 0
	currentPos := startingPos
	prevPos := [2]int{-1, -1}
	newStart := replaceStart(startingPos, graph)

	graph[startingPos[0]][startingPos[1]] = newStart
	
    fmt.Println("New start", newStart)
	for different(startingPos, currentPos) || length == 0 {
		i, j := currentPos[0], currentPos[1]

        pos := [2]int{-1, -1}
		for {
            // if south check isNorth
			stepI := 1
			stepJ := 0
			pos = [2]int{i + stepI, j + stepJ}
            if isSouth(currentPos, graph) {
                if isValid(pos, graph) && isNorth(pos, graph) && different(pos, prevPos) {
                    break
                }
            }
            // if north check isSouth
			stepI = -1
			stepJ = 0
			pos = [2]int{i + stepI, j + stepJ}
            if isNorth(currentPos, graph) {
                if isValid(pos, graph) && isSouth(pos, graph) && different(pos, prevPos) {
                    break
                }
            }
            // if east check isWest
			stepI = 0
			stepJ = 1
			pos = [2]int{i + stepI, j + stepJ}
            if isEast(currentPos, graph) {
                if isValid(pos, graph) && isWest(pos, graph) && different(pos, prevPos) {
                    break
                }
            }
            // if west check isEast
			stepI = 0
			stepJ = -1
			pos = [2]int{i + stepI, j + stepJ}
            if isWest(currentPos, graph) {
                if isValid(pos, graph) && isEast(pos, graph) && different(pos, prevPos) {
                    break
                }
            }
            panic("Failed to find")
		}

		prevPos = currentPos
		currentPos = pos
		length++
		fmt.Println(length, currentPos)
        //fmt.Print("\r", length)
	}
	return length
}

func main() {
	data, startingPos := parsedata("data/day10_input.txt")
	length := solve(data, startingPos)
    fmt.Println()
	fmt.Println(length / 2)
}
