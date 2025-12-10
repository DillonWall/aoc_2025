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

var result int64 = 0

func min(i1, i2 int) int {
    if i1 < i2 {
        return i1
    }
    return i2
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	file, err := os.Open("test.txt")
	// file, err := os.Open("input.txt")
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

	fmt.Printf("Result is %v\n", result)
}

func processText(lineno int, line string) {
	parts := strings.Split(line, " ")
	goal := 0
	buttons := make([][]int, 0)
	// joltages := make([]int, 0)
	for i := 1; i < len(parts[0]) - 1; i++ {
		if parts[0][i] == '#' {
			goal |= 1 << (i - 1)
		}
	}
	for i := 1; i < len(parts) - 1; i++ {
		numParts := strings.Split(parts[i][1:len(parts[i])-1], ",")
		nums := make([]int, len(numParts))
		for j, part := range numParts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Error: %v/n", err)
			}
			nums[j] = num
		}
		buttons = append(buttons, nums)
	}

	visited := make(map[int]bool)
	queue := []int{0}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		if visited[state] {
			continue
		}
		visited[state] = true

		current := 0
		for b := 0; b < len(buttons); b++ {
			if (state & (1 << b)) == 1 {
				for _, pos := range buttons[b] {
					// toggle pos in current
					current ^= (1 << pos)
				}
			}
		}
		// TODO
	}
}

func equalBoolSlices(a, b []bool) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
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
