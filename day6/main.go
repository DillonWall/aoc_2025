package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	Nums     []string
	Operator string
}

var result int64 = 0
var problems []Problem = make([]Problem, 0)
var lines []string = make([]string, 0)

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

    processNums()

	fmt.Printf("Result is %v\n", result)
}

func processText(line string) {
	lines = append(lines, line)
}

func processNums() {
	// add all nums and operators from lines to problems
	colStrs := make([]string, 0)
    operator := ""
	for i := 0; i < len(lines[0]) + 5; i++ {
		allSpace := true
        colStrs = append(colStrs, "")
		for j := 0; j < len(lines) - 1; j++ {
            if i >= len(lines[j]) {
                continue
            }
            newChar := lines[j][i]
			colStrs[len(colStrs)-1] += string(newChar)
			if newChar != ' ' {
				allSpace = false
			}
		}
        if i < len(lines[len(lines)-1]) {
            operator += lines[len(lines)-1][i:i+1]
        }

        if allSpace && operator == "" {
            break
        }
		if allSpace {
            fmt.Printf("ColStrs: %+v with operator %s\n", colStrs, operator)
			problems = append(problems, Problem{Nums: colStrs[:len(colStrs)-1], Operator: operator})
            colStrs = make([]string, 0)
            operator = ""
		}
	}

    fmt.Printf("Problems: %+v\n", problems)

	for _, problem := range problems {
		op := strings.Trim(problem.Operator, " ")
		res := 0
		if op == "+" {
			for _, n := range problem.Nums {
				nInt, err := strconv.Atoi(strings.Trim(n, " "))
				if err != nil {
					fmt.Printf("Error converting string to int: %v\n", err)
				}
				res += nInt
			}
		} else if op == "*" {
			res = 1
			for _, n := range problem.Nums {
				nInt, err := strconv.Atoi(strings.Trim(n, " "))
				if err != nil {
					fmt.Printf("Error converting string to int: %v\n", err)
				}
				res *= nInt
			}
		}
		result += int64(res)
	}
}
