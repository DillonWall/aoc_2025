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

var result uint64 = 0
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

type State struct {
    Node string
    Dac bool
    Fft bool
}

func calculateResult() {
    cache := make(map[State]uint64)

    var dfs func(State) uint64
    dfs = func(state State) uint64 {
        if val, exists := cache[state]; exists {
            return val
        }

        if state.Node == "out" {
            if state.Dac && state.Fft {
                cache[state] = 1
            } else { cache[state] = 0 }
            return cache[state]
        }

        sum := uint64(0)
        for _, next := range devices[state.Node] {
            nextState := State{next, state.Dac, state.Fft}
            if next == "dac" { nextState.Dac = true }
            if next == "fft" { nextState.Fft = true }
            sum += dfs(nextState)
        }
        cache[state] = sum
        return sum
	}

    result = dfs(State{"svr", false, false})
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
