package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
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

	buf, _ := os.ReadFile("day01/" + filename)

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
	list1 := []int{}
	list2 := []int{}

	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		values := strings.Split(s.Text(), "   ")
		if i, err := strconv.Atoi(values[0]); err == nil {
			list1 = append(list1, i)
		}
		if i, err := strconv.Atoi(values[1]); err == nil {
			list2 = append(list2, i)
		}
	}
	sort.Ints(list1)
	sort.Ints(list2)

	result := 0
	locationZip, _ := zip(list1, list2)

	for _, tuple := range locationZip {
		result = result + diff(tuple.a, tuple.b)
	}

	return result, true
}

func part2(ctx Context, b []byte) (int, bool) {
	if ctx.part != 2 {
		return 0, false
	}

	locationCount := map[int]intTuple{}
	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		values := strings.Split(s.Text(), "   ")
		location1, _ := strconv.Atoi(values[0])
		locationCount[location1] = intTuple{locationCount[location1].a + 1, locationCount[location1].b}

		location2, _ := strconv.Atoi(values[1])
		locationCount[location2] = intTuple{locationCount[location2].a, locationCount[location2].b + 1}
	}

	result := 0
	for k, v := range locationCount {
		result = result + k*v.a*v.b
	}

	return result, true
}
