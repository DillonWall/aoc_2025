package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Lower int64
	Upper int64
}

var result int64 = 0
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
		if !readingRange {
			calculateIds()
			break
		}
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

func max(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}

func calculateIds() {
	// merge ranges
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Lower < ranges[j].Lower
	})
	fmt.Printf("sorted ranges: %v\n", ranges)
	newRanges := ranges[:0]
	prev := ranges[0]
	for i := 1; i < len(ranges); i++ {
		if ranges[i].Lower <= prev.Upper {
			prev = Range{prev.Lower, max(ranges[i].Upper, prev.Upper)}
		} else {
			newRanges = append(newRanges, prev)
			prev = ranges[i]
		}
	}
	newRanges = append(newRanges, prev)
	ranges = newRanges
	fmt.Printf("merged ranges: %v\n", ranges)


	for _, r := range ranges {
		result += r.Upper - r.Lower + 1
		fmt.Printf("adding %v\n", r.Upper - r.Lower + 1)
	}
}
