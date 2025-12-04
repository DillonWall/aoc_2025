package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var result = 0

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

	fmt.Printf("Result is %v\n", result)
}

func processText(line string) {
	// find best first spot
	prevPos := len(line) - 12 // start 12th from end
	for i := prevPos - 1; i >= 0; i-- {
		if line[i] >= line[prevPos] {
			prevPos = i
		}
	}

	// find next 11 best spots
	chrs := make([]byte, 12)
	chrs[0] = line[prevPos]
	for spot := range 11 {
		nextPos := len(line) - (11 - spot) // start 11th, 10th, etc. from the end
		for i := nextPos - 1; i > prevPos; i-- {
			if line[i] >= line[nextPos] {
				nextPos = i
			}
		}
		chrs[spot+1] = line[nextPos]
		prevPos = nextPos
	}

	res, err := strconv.Atoi(string(chrs))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("adding %v\n", res)
	result += res
}
