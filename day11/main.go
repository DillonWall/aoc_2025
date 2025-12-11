package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	X int
	Y int
}

var result int64 = 0
var devices map[string][]string = make(map[string][]string)

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

func processText(lineno int, line string) {
	parts := strings.Split(line, " ")
    devices[parts[0][0:3]] = parts[1:]
}

func calculateResult() {
	queue := [][]string{{"you"}}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

        if state[len(state)-1] == "out" {
            result++
            continue
        }

        for _, d := range devices[state[len(state)-1]] {
            nextState := append([]string{}, state...)
            nextState = append(nextState, d)
            queue = append(queue, nextState)
        }
	}
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
