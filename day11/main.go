package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	X int
	Y int
}

var result int64 = 0
var devices map[string][]string = make(map[string][]string)
var inputs map[string][]string = make(map[string][]string)

func main() {
	file, err := os.Open("test.txt")
	// file, err := os.Open("input.txt")
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

    calculateResult()

	fmt.Printf("Result is %v\n", result)
}

func processText(lineno int, line string) {
	parts := strings.Split(line, " ")
    devices[parts[0][0:3]] = parts[1:]
	inputs[parts[0][0:3]] = []string{}
}

func calculateResult() {
	for k, v := range devices {
		for _, out := range v {
			inputs[out] = append(inputs[out], k)
		}
	}

	dac := numPathsFromAToB("svr", "dac")
	dacToFft := numPathsFromAToB("dac", "fft")
	fftToOut := numPathsFromAToB("fft", "out")

	fft := numPathsFromAToB("svr", "fft")
	fftToDac := numPathsFromAToB("fft", "dac")
	dacToOut := numPathsFromAToB("dac", "out")

	fmt.Printf("%v %v %v %v %v %v\n", dac, dacToFft, fftToOut, fft, fftToDac, dacToOut)

	result = int64((dac * dacToFft * fftToOut) + (fft * fftToDac * dacToOut))
}

func numPathsFromAToB(a, b string) int {
	fmt.Printf("Calculating paths from %v to %v\n", a, b)
	memo := map[string]int{a: 1}
	queue := []string{a}
	done := map[string]struct{}{a: {}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Printf("Processing %v\n", cur)
		for _, out := range devices[cur] {
			memo[out] += memo[cur]
			allDone := true
			fmt.Printf("Checking %v inputs: %v\n", out, inputs[out])
			for _, in := range inputs[out] {
				if _, exists := done[in]; !exists {
					fmt.Printf("  Not done: %v\n", in)
					allDone = false
				}
			}

			if out == b {
				continue
			}
			if allDone {
				fmt.Printf("  Done: %v\n", out)
				done[out] = struct{}{}
				queue = append(queue, out)
			}
		}
	}

	fmt.Printf("Paths from %v to %v: %v\n", a, b, memo[b])
	fmt.Printf("final counted paths to %v: %v\n", b, memo[b])
	return memo[b]
}

func calculateResultOld() {
	stack := [][]string{{"svr"}}
	memo := []string{"out", "dac", "fft"}
	count := 0
	reachedOut := 0
	dacs, ffts := 0, 0
	for len(stack) > 0 {
		count++
		state := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack) > 5000000 {
			panic("huge stack")
		}
		if count % 1000000 == 0 {
			fmt.Printf("count: %v, result: %v, reachedOut: %v, dacs: %v, ffts: %v\n", count, result, reachedOut, dacs, ffts)
		}

		for _, v := range memo {
			if v == state[len(state)-1] {
				reachedOut++
				dac, fft := false, false
				for _, s := range state {
					if s == "dac" { dac = true }
					if s == "fft" { fft = true }
				}
				if dac { dacs++ }
				if fft { ffts++ }
				if dac && fft {
					for _, s := range state {
						if s == "dac" || s == "fft" {
							break
						}
						memo = append(memo, s)
					}
					result++
				}
				continue
			}
		}

        for _, d := range devices[state[len(state)-1]] {
            nextState := append([]string{}, state...)
            nextState = append(nextState, d)
            stack = append(stack, nextState)
        }
	}
}

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
