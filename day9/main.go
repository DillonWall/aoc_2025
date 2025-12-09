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
    // get lines between all coords, including last to first
    lines := make([][2]Coord, 0)
    for i := 0; i < len(coords); i++ {
        j := (i + 1) % len(coords)
        lines = append(lines, [2]Coord{coords[i], coords[j]})
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

	for i := len(areas)-1; i >= 0; i-- {
		a := areas[i]
        aMinX := min(a.P1.X, a.P2.X)
        aMaxX := max(a.P1.X, a.P2.X)
        aMinY := min(a.P1.Y, a.P2.Y)
        aMaxY := max(a.P1.Y, a.P2.Y)
        // check if any line intersects with area
        found := false
        for _, line := range lines {
            minX := min(line[0].X, line[1].X)
            maxX := max(line[0].X, line[1].X)
            minY := min(line[0].Y, line[1].Y)
            maxY := max(line[0].Y, line[1].Y)
            vertIntersects := (minX < aMaxX && maxX > aMinX) && !(minY > aMaxY) && !(maxY < aMinY)
            horzIntersects := (minY < aMaxY && maxY > aMinY) && !(minX > aMaxX) && !(maxX < aMinX)
            if minX == maxX {

            if vertIntersects || horzIntersects {
                fmt.Printf("Area %v between points (%v,%v) and (%v,%v) intersects with line between (%v,%v) and (%v,%v)\n", a.A, a.P1.X, a.P1.Y, a.P2.X, a.P2.Y, line[0].X, line[0].Y, line[1].X, line[1].Y)
                found = true
                break
            }
        }
		if !found {
            fmt.Printf("Found area %v between points (%v,%v) and (%v,%v)\n", a.A, a.P1.X, a.P1.Y, a.P2.X, a.P2.Y)
			result = int64(a.A)
			return
		}
	}
}
