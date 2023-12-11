package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type MoveSequence struct {
    Moves []string
}

type Cycler struct {
    Seq *MoveSequence
    Step int
}

func (c *Cycler) Cycle() string {
    ret := c.Seq.Moves[c.Step]
    c.Step += 1
    c.Step = c.Step % len(c.Seq.Moves)
    return ret
}

type Node struct {
    Left, Right string
}
type Graph map[string]Node

func step(instruction, current string, graph Graph) (next string) {
    if instruction == "R" {
        return graph[current].Right
    }
    return graph[current].Left
}

func walkThrough(instructions MoveSequence, graph Graph, start string) (steps int) {
    current := start
    
    cycler := Cycler{&instructions, 0}

    for ; current[2] != 'Z'; steps++ {
        instruction := cycler.Cycle()
        current = step(instruction, current, graph)
    }

    return steps
}

func buildGraph(lines []string) (graph Graph) {
    nodeRe := regexp.MustCompile(`(?P<node>\w{3}) = \((?P<left>\w{3}), (?P<right>\w{3})\)`)
    nodeIndex := nodeRe.SubexpIndex("node")
    leftIndex := nodeRe.SubexpIndex("left")
    rightIndex := nodeRe.SubexpIndex("right")

    graph = make(Graph)
    for _, line := range lines {
        if len(line) == 0 {
            continue
        }
        parsed := nodeRe.FindStringSubmatch(line)
        nodeName := parsed[nodeIndex]
        leftName := parsed[leftIndex]
        rightName := parsed[rightIndex]

        node := Node{leftName, rightName}
        graph[nodeName] = node
    }

    return graph
}

func parseData(filename string) (ms MoveSequence, graph Graph) {
    filecontent, err := os.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    lines := strings.Split(string(filecontent), "\n")
    ms = MoveSequence{strings.Split(lines[0], "")}

    graph = buildGraph(lines[2:])

    return ms, graph
}

func findStarts(graph *Graph) (starts []string) {
    for k := range *graph {
        if k[2] == 'A' {
            starts = append(starts, k)
        }
    }
    return starts
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

func lcm(a, b int) int {
    result := a * b / gcd(a, b)
    return result
}

func lcmSlice(ints []int) int {
    ret := lcm(ints[0], ints[1])
    for _, i := range ints[2:] {
        ret = lcm(ret, i)
    }
    return ret
}

func main() {
    moveSequence, graph := parseData("data/day8_input.txt")
    starts := findStarts(&graph)
    var steps []int
    for _, start := range starts {
        nbStep := walkThrough(moveSequence, graph, start)
        steps = append(steps, nbStep)
    }
    totalSteps := lcmSlice(steps)
    fmt.Println(totalSteps)
}
