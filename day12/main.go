package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

type Problem struct {
    Width int
    Height int
    NumPresent []int
}

var result uint64 = 0
var options [][]string = make([][]string, 6)
var problems []Problem = make([]Problem, 0)

func main() {
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    i := 0
    for scanner.Scan() {
		line := scanner.Text()
		processText(i, line)
        i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

    calculateResult()

	fmt.Printf("Result is %v\n", result)
}

var currentOption int = 0
func processText(lineno int, line string) {
    if len(line) == 0 {
        currentOption++
        if currentOption < 6 {
            options[currentOption] = make([]string, 0)
        }
        return
    }

    if line[0] == '#' || line[0] == '.' {
        options[currentOption] = append(options[currentOption], line)
    }

    if currentOption > 5 {
        parts := strings.Split(line, " ")
        whParts := strings.Split(parts[0][:len(parts[0])-1], "x")
        width, err := strconv.Atoi(whParts[0])
        if err != nil {
            panic(err)
        }
        height, err := strconv.Atoi(whParts[1])
        if err != nil {
            panic(err)
        }
        numPresent := make([]int, 6)
        for i := 0; i < 6; i++ {
            numPresent[i], err = strconv.Atoi(parts[i+1])
            if err != nil {
                panic(err)
            }
        }
        problem := Problem{width, height, numPresent}
        problems = append(problems, problem)
    }
}

type State struct {
    Node string
    Dac bool
    Fft bool
}

func calculateResult() {
    sizes := make([]int, 6)
    for i := 0; i < 6; i++ {
        count := 0
        for _, line := range options[i] {
            for _, ch := range line {
                if ch == '#' {
                    count++
                }
            }
        }
        sizes[i] = count
    }

    fmt.Printf("Sizes: %v\n", sizes)

    numProblems := len(problems)
    numPossible := 0
    for _, problem := range problems {
        // check if possible
        possible := true
        tilesLowerBound := 0
        availableSpace := problem.Width * problem.Height
        for i := 0; i < 6; i++ {
            tilesLowerBound += problem.NumPresent[i] * sizes[i]
        }
        if tilesLowerBound > availableSpace {
            possible = false
        }
        if possible {
            numPossible++
        }
    }

    fmt.Printf("num possible out of total: %v / %v\n", numPossible, numProblems)
}

func Min(i1, i2 int) int {
    if i1 < i2 {
        return i1
    }
    return i2
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type Pair struct {
    Key   int
    Value int
}

type PairHeap []Pair

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].Value > h[j].Value } // Max-heap by Value
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PairHeap) Push(x any) { *h = append(*h, x.(Pair)) }
func (h *PairHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
