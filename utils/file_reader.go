package utils

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
)

func ReadFileToIntSlice(f *os.File) []int {
	scanner := bufio.NewScanner(f)
	slice := make([]int, 0, 100)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		slice = append(slice, num)
	}

	return slice
}

func ReadFileToStringSlice(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	slice := make([]string, 0, 100)

	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	return slice
}

func ReadFileToStringList(f *os.File) *list.List {
	scanner := bufio.NewScanner(f)
	input := list.New()

	for scanner.Scan() {
		input.PushBack(scanner.Text())
	}

	return input
}
