package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type intTuple struct {
	a, b int
}

// Context defined which part to run
type Context struct {
	part int
}

func zip(a, b []int) ([]intTuple, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}

	r := make([]intTuple, len(a))

	for i, e := range a {
		r[i] = intTuple{e, b[i]}
	}

	return r, nil
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getDir(a int) int {
	if a >= 0 {
		return 1
	}
	return -1
}

func main() {
	var part int
	var input bool
	flag.IntVar(&part, "p", 1, "part 1 or 2")
	flag.BoolVar(&input, "e", false, "use example.txt instead of input.txt")
	flag.Parse()

	filename := "input.txt"
	if input {
		filename = "example.txt"
	}

	fmt.Println("Running part", part, "with", filename)

	buf, _ := os.ReadFile("day02/" + filename)

	ctx := Context{
		part: part,
	}

	partHandlers := []func(Context, []byte) (int, bool){part1, part2}
	for _, handler := range partHandlers {
		if result, ok := handler(ctx, buf); ok {
			fmt.Println(result)
		}
	}
}

func part1(ctx Context, b []byte) (int, bool) {
	if ctx.part != 1 {
		return 0, false
	}

	result := 0

	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		safe := true
		dir := 0
		report := strings.Split(s.Text(), " ")
		for i, lvlStr := range report[:len(report)-1] {
			if !safe {
				continue
			}

			currentLvl, _ := strconv.Atoi(lvlStr)
			nextLvl, _ := strconv.Atoi(report[i+1])

			if diff(currentLvl, nextLvl) > 3 {
				safe = false
			}

			if currentLvl == nextLvl {
				safe = false
			}

			if dir != 0 && dir != getDir(currentLvl-nextLvl) {
				safe = false
			}

			dir = getDir(currentLvl - nextLvl)
		}
		if safe {
			result++
		}
	}

	return result, true
}

func isSafe(dir, a, b int) bool {
	if diff(a, b) > 3 {
		return false
	}

	if a == b {
		return false
	}

	if dir != 0 && dir != getDir(a-b) {
		return false
	}

	return true
}

func part2(ctx Context, b []byte) (int, bool) {
	if ctx.part != 2 {
		return 0, false
	}

	result := 0

	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		safe := true
		dir := 0
		report := strings.Split(s.Text(), " ")
		dampenerUsed := false
		for i := 0; i < len(report)-1; i++ {
			if !safe {
				continue
			}

			currentLvl, _ := strconv.Atoi(report[i])
			nextLvl, _ := strconv.Atoi(report[i+1])

			safe = isSafe(dir, currentLvl, nextLvl)
			if safe {
				dir = getDir(currentLvl - nextLvl)
			}

			if !safe && !dampenerUsed {
				safe = true
				dampenerUsed = true

				if i+2 == len(report) {
					continue
				}

				if i == 0 {
					nextNextLvl, _ := strconv.Atoi(report[i+2])
					if isSafe(dir, currentLvl, nextNextLvl) {
						dir = getDir(currentLvl - nextNextLvl)
						i++
					} else if isSafe(dir, nextLvl, nextNextLvl) {
						dir = getDir(nextLvl - nextNextLvl)
						i++
					} else {
						safe = false
					}
					continue
				}

				nextNextLvl, _ := strconv.Atoi(report[i+2])
				previousLvl, _ := strconv.Atoi(report[i-1])
				if isSafe(dir, previousLvl, nextLvl) {
					dir = getDir(previousLvl - nextLvl)
					i++
				} else if isSafe(dir, currentLvl, nextNextLvl) {
					dir = getDir(currentLvl - nextNextLvl)
					i++
				} else {
					safe = false
				}
			}
		}
		if safe {
			result++
		}
	}

	return result, true
}
