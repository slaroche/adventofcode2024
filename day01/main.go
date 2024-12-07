package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type intTuple struct {
	a, b int
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
	filename := "input.txt"
	buf, _ := os.ReadFile("day01/" + filename)

	// Part 1
	s := bufio.NewScanner(bytes.NewReader(buf))

	list1 := []int{}
	list2 := []int{}

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

	result1 := 0
	zip1, _ := zip(list1, list2)

	for _, tuple := range zip1 {
		result1 = result1 + diff(tuple.a, tuple.b)
	}
	fmt.Println(result1)

	// Part 2
	s = bufio.NewScanner(bytes.NewReader(buf))

	locationCount := map[int]intTuple{}
	for s.Scan() {
		values := strings.Split(s.Text(), "   ")
		location1, _ := strconv.Atoi(values[0])
		locationCount[location1] = intTuple{locationCount[location1].a + 1, locationCount[location1].b}

		location2, _ := strconv.Atoi(values[1])
		locationCount[location2] = intTuple{locationCount[location2].a, locationCount[location2].b + 1}
	}

	result2 := 0
	for k, v := range locationCount {
		result2 = result2 + k*v.a*v.b
	}

	fmt.Println(result2)
}
