package main

import (
	"bufio"
	"fmt"
	"os"
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

    calculateResults()

	fmt.Printf("Result is %v\n", result)
}

func processText(line string) {
    // if line contains + or * then its a op line
    operationLine := false
    if strings.Contains(line, "+") || strings.Contains(line, "*") {
        operationLine = true
    }
    firstLine := false
    if problems == nil || len(problems) == 0 {
        firstLine = true
    }

    parts := strings.Fields(line)
    if !operationLine {
        for i, part := range parts {
            num, err := strconv.Atoi(part)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            }
            if firstLine {
                problems = append(problems, Problem{Nums: []int{num}})
            } else {
                problems[i].Nums = append(problems[i].Nums, num)
            }
        }
    } else {
        for i, part := range parts {
            problems[i].Operator = part[0]
            res := 0
            if problems[i].Operator == '+' {
                for _, n := range problems[i].Nums {
                    res += n
                }
            } else if problems[i].Operator == '*' {
                res = 1
                for _, n := range problems[i].Nums {
                    res *= n
                }
            }
            result += int64(res)
        }
    }
}

func calculateResults() {
}
