package utils

import (
	"errors"
	"os"
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
