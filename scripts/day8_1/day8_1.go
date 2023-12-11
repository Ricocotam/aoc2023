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

func walkThrough(instructions MoveSequence, graph Graph) (steps int) {
    current := "AAA"
    end := "ZZZ"
    
    cycler := Cycler{&instructions, 0}

    for ; current != end; steps++ {
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

func main() {
    moveSequence, graph := parseData("data/day8_input.txt")
    fmt.Println("Starting")
    steps := walkThrough(moveSequence, graph)
    fmt.Println(steps)
}
