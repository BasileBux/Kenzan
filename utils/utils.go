package utils

import (
	"errors"
	"os"
	"strings"
)

// utils package is a collection of missing stuff in golang

func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func PushFront(x []int, y int) []int {
	x = append([]int{y}, x...)
	return x
}

// solutions from: https://stackoverflow.com/a/12518877
func FileExists(filename string) bool {
	if _, err := os.Stat("/path/to/whatever"); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		// We don't know ...
		return false
	}
}

func FindAllIndices(str, substr string) []int {
	var indices []int
	start := 0
	for {
		index := strings.Index(str[start:], substr)
		if index == -1 {
			break
		}
		indices = append(indices, start+index)
		start += index + 1
	}
	return indices
}
