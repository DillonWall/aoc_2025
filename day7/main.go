package main

import (
	"bufio"
	"fmt"
	"os"
)

var result int64 = 0
var beams = make([]int, 0)

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
    if lineno == 0 {
        for _, c := range line {
            if c == 'S' {
                beams = append(beams, 1)
            } else {
                beams = append(beams, 0)
            }
        }
        return
    }

    for i, c := range line {
        if c == '^' && beams[i] > 0 {
            if i > 0 {
                beams[i-1] += beams[i]
            }
            if i < len(beams)-1 {
                beams[i+1] += beams[i]
            }
            beams[i] = 0
        }
    }
}

func calculateResults() {
    for _, v := range beams {
        result += int64(v)
    }
}
