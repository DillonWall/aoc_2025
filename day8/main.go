package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
	Z int
}

const NUM_BOXES = 1000
var result int64 = 0
var coords []Coord = make([]Coord, 0)

func max(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}

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

    calculateResults()

	fmt.Printf("Result is %v\n", result)
}

func processText(lineno int, line string) {
	parts := strings.Split(line, ",")
	nums := make([]int, 3)
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Error: %v/n", err)
		}
		nums[i] = num
	}
	coords = append(coords, Coord{nums[0], nums[1], nums[2]})
}

func dsquared(point1, point2 Coord) int {
	return (point1.X - point2.X) * (point1.X - point2.X) +
		   (point1.Y - point2.Y) * (point1.Y - point2.Y) +
		   (point1.Z - point2.Z) * (point1.Z - point2.Z)
}

type Dist struct {
	D int
	P1 int
	P2 int
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

func calculateResults() {
	dists := make([]Dist, 0)
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			dists = append(dists, Dist{dsquared(coords[i], coords[j]), i, j})
		}
    }
	fmt.Printf("dists: %v\n", dists)

	sort.Slice(dists, func(i, j int) bool {
		return dists[i].D < dists[j].D
	})

	cToJunc := make(map[int]int)
	curJunc := 1
	for _, dist := range dists[:NUM_BOXES] {
		j1, ok1 := cToJunc[dist.P1]
		j2, ok2 := cToJunc[dist.P2]
		if !ok1 && !ok2 {
			cToJunc[dist.P1] = curJunc
			cToJunc[dist.P2] = curJunc
			curJunc++
		} else if !ok1 {
			cToJunc[dist.P1] = j2
		} else if !ok2 {
			cToJunc[dist.P2] = j1
		} else {
			for k, v := range cToJunc {
				if v == j2 {
					cToJunc[k] = j1
				}
			}
		}
	}

	counts := make(map[int]int, 0)
	for _, v := range cToJunc {
		counts[v]++
	}

	pairs := make(PairHeap, 0, len(counts))
	for k, v := range counts {
		pairs = append(pairs, Pair{Key: k, Value: v})
	}

	heap.Init(&pairs)

	result = 1
	result *= int64(heap.Pop(&pairs).(Pair).Value)
	result *= int64(heap.Pop(&pairs).(Pair).Value)
	result *= int64(heap.Pop(&pairs).(Pair).Value)
}
