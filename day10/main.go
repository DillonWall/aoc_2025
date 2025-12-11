package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
)

type Coord struct {
	X int
	Y int
}

var result int64 = 0

func Min(i1, i2 int) int {
    if i1 < i2 {
        return i1
    }
    return i2
}

func Max(i, j int) int {
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

	fmt.Printf("Result is %v\n", result)
}

func processText(lineno int, line string) {
	parts := strings.Split(line, " ")
	buttons := make([][]int, 0)
	jStr := parts[len(parts)-1]
	joltageStrs := strings.Split(jStr[1:len(jStr)-1], ",")
	joltages := make([]int, len(joltageStrs))
	for i, js := range joltageStrs {
		num, err := strconv.Atoi(js)
		if err != nil {
			fmt.Printf("Error: %v/n", err)
		}
		joltages[i] = num
	}
	for i := 1; i < len(parts) - 1; i++ {
		numParts := strings.Split(parts[i][1:len(parts[i])-1], ",")
		nums := make([]int, len(numParts))
		for j, part := range numParts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Error: %v/n", err)
			}
			nums[j] = num
		}
		buttons = append(buttons, nums)
	}
	buttonEffects := make([][]int, len(buttons))
	for i := 0; i < len(buttons); i++ {
		buttonEffects[i] = make([]int, len(joltages))
		for _, j := range buttons[i] {
			buttonEffects[i][j] = 1
		}
	}
	solveWithLP(joltages, buttonEffects)
}

func solveWithLP(target []int, buttonEffects [][]int) {
    m := len(buttonEffects) // number of buttons
    n := len(target)        // number of joltages
    fmt.Printf("Solving LP with %d buttons and %d joltages\n", m, n)
    fmt.Printf("Button effects: %v\n", buttonEffects)
    fmt.Printf("Target: %v\n", target)

    lp := golp.NewLP(0, m)

    // Add equality constraints: sum(buttonEffects[i][j] * x_i) == target[j]
    for j := 0; j < n; j++ {
        coeffs := make([]float64, m)
        for i := 0; i < m; i++ {
            coeffs[i] = float64(buttonEffects[i][j])
        }
        lp.AddConstraint(coeffs, golp.EQ, float64(target[j]))
    }

    // Objective: minimize total presses (default is minimization)
    obj := make([]float64, m)
    for i := 0; i < m; i++ {
        obj[i] = 1.0
    }
    lp.SetObjFn(obj)

    // Set variables to be integer and >= 0
    for i := 0; i < m; i++ {
        lp.SetInt(i, true)
        lp.SetBounds(i, 0.0, math.Inf(1))
    }

    lpresult := lp.Solve()
    if lpresult != golp.OPTIMAL {
        fmt.Println("No solution found")
        return
    }

    vars := lp.Variables()
	count := 0
    for i := 0; i < m; i++ {
        fmt.Printf("Button %d: %.0f times\n", i, vars[i])
		count += int(vars[i])
    }

	result += int64(count)
}

type Pair struct {
    Key   int
    Value int
}

type PairHeap []Pair

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].Value > h[j].Value } // Max-heap by Value
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PairHeap) Push(x any) { *h = append(*h, x.(Pair)) }
func (h *PairHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
