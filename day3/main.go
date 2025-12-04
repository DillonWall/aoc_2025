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
	highestPos := len(line) - 2 // start 2nd from end
	// highest := int(line[highestPos] - '0')

	// find best first spot
	for i := highestPos - 1; i >= 0; i-- {
		if line[i] >= line[highestPos] {
			highestPos = i
			// highest = int(line[highestPos] - '0')
		}
	}

	// find best second spot
	rPos := len(line) - 1 // start at end
	// rHigh := int(line[rPos] - '0')
	for i := rPos - 1; i > highestPos; i-- {
		if line[i] >= line[rPos] {
			rPos = i
			// rHigh = int(line[rPos] - '0')
		}
	}

	res, err := strconv.Atoi(string([]byte{line[highestPos], line[rPos]}))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("adding %v\n", res)
	result += res
}
