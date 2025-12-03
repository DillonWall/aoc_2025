package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var currentNum = 50
var password = 0

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

	fmt.Printf("Password is %v\n", password)
}

func processText(line string) {
	letter := line[:1]
	numStr := line[1:]

	dir := 1
	if letter == "L" {
		dir = -1
	}
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// flip direction for negatives
	if num < 0 {
		dir = dir * -1
		num = -num
	}

	currentNum += (num % 100) * dir
	if currentNum < 0 {
		currentNum += 100
	}
	if currentNum > 99 {
		currentNum -= 100
	}

	fmt.Printf("currentNum = %v\n", currentNum)
	if currentNum == 0 {
		password++
	}
}
