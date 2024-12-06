package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {
	buf, err := os.ReadFile("day01/input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(bytes.NewReader(buf))

	array := []string{}

	for s.Scan() {
		array = append(array, s.Text())
	}
	for _, line := range array {
		println(line)
	}
}
