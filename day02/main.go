package main

import (
	"adventofcode2024/utils"
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func getDir(a int) int {
	if a >= 0 {
		return 1
	}
	return -1
}

func isSafe(report []int) (bool, int) {
	safe := true
	dir := 0
	index := 0
	for i, lvl := range report[:len(report)-1] {
		if !safe {
			continue
		}

		nextLvl := report[i+1]

		if utils.Diff(lvl, nextLvl) > 3 {
			safe = false
		}

		if lvl == nextLvl {
			safe = false
		}

		if dir != 0 && dir != getDir(lvl-nextLvl) {
			safe = false
		}

		dir = getDir(lvl - nextLvl)
		index = i
	}
	return safe, index
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

	buf, err := os.ReadFile("day02/" + filename)
	if err != nil {
		log.Fatal("cannot load file")
	}

	ctx := utils.Context{
		Part: part,
	}

	partHandlers := []func(utils.Context, []byte) (int, bool){part1, part2}
	for _, handler := range partHandlers {
		if result, ok := handler(ctx, buf); ok {
			fmt.Println(result)
		}
	}
}

func part1(ctx utils.Context, b []byte) (int, bool) {
	if ctx.Part != 1 {
		return 0, false
	}

	result := 0

	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		strReport := strings.Split(s.Text(), " ")
		report, _ := utils.StrSliceToInt(strReport)
		if safe, _ := isSafe(report); safe {
			result++
		}
	}

	return result, true
}

func part2(ctx utils.Context, b []byte) (int, bool) {
	if ctx.Part != 2 {
		return 0, false
	}

	result := 0

	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		safe := true
		strReport := strings.Split(s.Text(), " ")
		report, _ := utils.StrSliceToInt(strReport)
		safe, index := isSafe(report)
		if safe {
			result++
		} else {
			if safe, _ = isSafe(utils.Remove(report, index)); safe {
				result++
			} else if safe, _ = isSafe(utils.Remove(report, index+1)); safe {
				result++
			}
		}
	}
	err := s.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result, true
}
