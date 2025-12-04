package main

import (
	"bufio"
	"fmt"
	"os"
)

var result = 0
var grid [][]bool

func main() {
	// file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		processText(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	lastResult := -1
	for result != lastResult {
		lastResult = result
		processGrid()
	}

	fmt.Printf("Result is %v\n", result)
}

func processText(line string) {
	if grid == nil {
		grid = make([][]bool, 0)
	}
	row := make([]bool, len(line))
	for i, chr := range line {
		row[i] = chr == '@'
	}
	grid = append(grid, row)
}

type Coord struct {
	Row int
	Col int
}

func processGrid() {
	toRemove := make([]Coord, 0)
	for row := range len(grid) {
		for col := range len(grid[0]) {
			if !grid[row][col] {
				continue
			}
			count := 0
			for dr := -1; dr <= 1; dr++ {
				if row + dr < 0 || row + dr >= len(grid) {
					continue
				}
				for dc := -1; dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					} else if col + dc < 0 || col + dc >= len(grid[0]) {
						continue
					} else if grid[row + dr][col + dc] {
						count++
					}
				}
			}
			if count < 4 {
				result++
				toRemove = append(toRemove, Coord{row, col})
			}
		}
	}

	for _, c := range toRemove {
		grid[c.Row][c.Col] = false
	}
}
