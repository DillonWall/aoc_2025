package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

var result int64 = 0
var coords []Coord = make([]Coord, 0)

func max(i, j int) int {
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

var lastX = 0
var lastY = 0

func processText(lineno int, line string) {
	parts := strings.Split(line, ",")
	nums := make([]int, 2)
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Error: %v/n", err)
		}
		nums[i] = num
	}
	coords = append(coords, Coord{nums[0], nums[1]})
	lastX = max(lastX, nums[0])
	lastY = max(lastY, nums[1])
}

func area(point1, point2 Coord) int {
	return int(math.Abs(float64(point1.X - point2.X) + 1) *
			   math.Abs(float64(point1.Y - point2.Y) + 1))
}

type Area struct {
	A int
	P1 Coord
	P2 Coord
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
	valid := make([][]bool, lastY+1)
	for i := range lastY+1 {
		valid[i] = make([]bool, lastX+1)
	}

	sign := func(x int) int {
		if x < 0 {
			return -1
		} else if x > 0 {
			return 1
		}
		return 0
	}

	for i := 0; i < len(coords); i++ {
		prev := coords[i]
		curr := coords[(i+1)%len(coords)] // wraps to first after last
		dx := sign(curr.X - prev.X)
		dy := sign(curr.Y - prev.Y)
		x, y := prev.X, prev.Y
		for x != curr.X || y != curr.Y {
			valid[y][x] = true
			x += dx
			y += dy
		}
	}

	// minX, maxX := coords[0].X, coords[0].X
	// minY, maxY := coords[0].Y, coords[0].Y
	// for _, c := range coords {
	// 	if c.X < minX { minX = c.X }
	// 	if c.X > maxX { maxX = c.X }
	// 	if c.Y < minY { minY = c.Y }
	// 	if c.Y > maxY { maxY = c.Y }
	// }
	// midX := (minX + maxX) / 2
	// midY := (minY + maxY) / 2
	midX := 96300
	midY := 56000

	queue := []Coord{Coord{midX, midY}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		x, y := p.X, p.Y
		if valid[y][x] {
			continue
		}
		valid[y][x] = true
		queue = append(queue, Coord{x + 1, y}, Coord{x - 1, y}, Coord{x, y + 1}, Coord{x, y - 1})
	}

	areas := make([]Area, 0)
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			areas = append(areas, Area{area(coords[i], coords[j]), coords[i], coords[j]})
		}
    }

	sort.Slice(areas, func(i, j int) bool {
		return areas[i].A < areas[j].A
	})

	// for _, arr := range valid {
	// 	fmt.Printf("%v\n", arr)
	// }

	for i := len(areas)-1; i >= 0; i-- {
		a := areas[i]
		done := false
		for x := min(a.P1.X, a.P2.X); x <= max(a.P1.X, a.P2.X); x++ {
			for y := min(a.P1.Y, a.P2.Y); y <= max(a.P1.Y, a.P2.Y); y++ {
				if !valid[y][x] {
					done = true
					break
				}
			}
			if done {
				break
			}
		}
		if !done {
			result = int64(a.A)
			return
		}
	}
}
