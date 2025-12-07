package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Problem struct {
	Nums []int
	Operator byte
}

var result int64 = 0
var problems []Problem = make([]Problem, 0)

func max(i, j int64) int64 {
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
	for scanner.Scan() {
		line := scanner.Text()
		processText(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

    calculateResults()

	fmt.Printf("Result is %v\n", result)
}

func processText(line string) {
    // if line contains + or * then its a op line

	if readingRange {
		split := strings.Split(line, "-")
		lower, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		upper, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		ranges = append(ranges, Range{lower, upper})
	} else {
		// id, err := strconv.Atoi(line)
		// if err != nil {
		// 	fmt.Printf("Error: %v\n", err)
		// }
		// checkId(id)
	}
}

func calculateResults() {
}
