package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Lower int
	Upper int
}

var result = 0
var readingRange = true
var ranges = make([]Range, 0)

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
	if line == "" {
		readingRange = false
		return
	}

	if readingRange {
		split := strings.Split(line, "-")
		lower, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		upper, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		ranges = append(ranges, Range{lower, upper})
	} else {
		id, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		checkId(id)
	}
}

func checkId(id int) {
	for _, r := range ranges {
		if id >= r.Lower && id <= r.Upper {
			result++
			return
		}
	}
}
