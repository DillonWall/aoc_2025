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
	fmt.Printf("%v\n", line)

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

	// account for negative meaning we passed 0 once
	if dir == -1 && num >= currentNum && currentNum != 0 {
		password++
		fmt.Println("inc password")
	}

	passwordInc := (currentNum + (num * dir)) / 100
	if passwordInc < 0 {
		passwordInc = -passwordInc
	}
	password += passwordInc
	if passwordInc > 0 {
		fmt.Printf("inc password by %v\n", passwordInc)
	}

	currentNum += (num * dir) % 100
	if currentNum < 0 {
		currentNum += 100
	} else if currentNum > 99 {
		currentNum -= 100
	}

	fmt.Printf("currentNum = %v\n", currentNum)
}
