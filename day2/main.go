package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	// "strconv"
	"strings"
)

var result = 0

func main() {
	// for n := 1; n < 16; n++ {
	//    fmt.Printf("Factors of %d: ", n)
	//    for i := 1; i <= n; i++ {
	//        if n%i == 0 {
	//            fmt.Printf("%d ", i)
	//        }
	//    }
	//    fmt.Println()
	// }

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
	ranges := strings.SplitSeq(line, ",")
	for rangeStr := range ranges {
		rangeParts := strings.Split(rangeStr, "-")
		processRange(rangeParts[0], rangeParts[1])
	}
}

var factors = map[int][]int {
	2: {1},
	3: {1},
	4: {1,2},
	5: {1},
	6: {1,2,3},
	7: {1},
	8: {1,2,4},
	9: {1,3},
	10: {1,2,5},
	11: {1},
	12: {1,2,3,4,6},
	13: {1},
	14: {1,2,7},
	15: {1,3,5},
}

// var factors = map[int][]int {
// 	2: {1},
// 	3: {},
// 	4: {2},
// 	5: {},
// 	6: {3},
// 	7: {},
// 	8: {4},
// 	9: {},
// 	10: {5},
// 	11: {},
// 	12: {6},
// 	13: {},
// 	14: {7},
// 	15: {},
// }

func processRange(lower string, upper string) {
	// 1234512340 - 1234512349
	// 123451234X
	// 1234512000 - 1234512999
	// 1234512XXX
	// 1111111XXX == 1
	// 1111111 15 23 44
	// 1 2 5
	// all 1
	// all 11
	// all 11111
	// 8XXXXXXX .. 8284583-8497825
	// 1 2 4
	// all 8
	// each iter spot is 8
	// all 8 X(2-4)
	// each iter spot is 82, 83, or 84
	// as we iterate, are we at the lowest or highest digits? if so, retain checks, else no checks
	// all 8 X(2-4) X(if 2 lower 8. if 4 upper 9) X(if 28 lower 4. if 49 upper 7)

	// handle different lower and upper len with recursion
	// if len(lower) < len(upper) {
	// 	newLower := "1" + strings.Repeat("0", len(lower))
	// 	processRange(newLower, upper)
	// 	upper = strings.Repeat("9", len(lower))
	// }
	// for _, factor := range factors[len(lower)] {
	// 	pattern := make([]byte, factor)
	// 	possible := 1
	// 	currLow := int(lower[0])
	// 	currHigh := int(upper[0])
	// 	for i := range factor {
	// 		for j := i; j < len(lower); j += factor {
	// 		}
	// 		// if lower[i] != upper[i] {
	// 		// 	possible += int(upper[i]) - int(lower[i])
	// 		// }
	// 	}
	// }

	// suboptimal approach:
	set := make(map[int]struct{})
	l, err := strconv.Atoi(lower)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	u, err := strconv.Atoi(upper)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	for i := l; i <= u; i++ {
		iStr := strconv.Itoa(i)
		for _, factor := range factors[len(iStr)] {
			shouldAdd := true
			for j := range factor {
				currDigit := iStr[j]
				for k := j; k < len(iStr); k += factor {
					if iStr[k] != currDigit {
						shouldAdd = false
						break
					}
				}
				if !shouldAdd {
					break
				}
			}
			if shouldAdd {
				set[i] = struct{}{}
			}
		}
	}
	for n := range set {
		fmt.Printf("adding %v\n", n)
		result += n
	}

	// id_int, err := strconv.Atoi(id)
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// }
	// result += id_int
}
