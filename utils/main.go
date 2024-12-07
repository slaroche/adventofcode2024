package utils

import (
	"fmt"
	"strconv"
)

// Context defined which part to run
type Context struct {
	Part int
}

// Diff return the different between 2 numbers
func Diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// StrSliceToInt convert a string slice to an int slice
func StrSliceToInt(s []string) ([]int, error) {
	intSlice := []int{}
	for _, lvl := range s {
		intLvl, err := strconv.Atoi(lvl)
		if err != nil {
			return nil, fmt.Errorf("Fails to convert string slice to int slice")
		}
		intSlice = append(intSlice, intLvl)
	}
	return intSlice, nil
}

// Remove is a function that remove en slice element at a given index
func Remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
