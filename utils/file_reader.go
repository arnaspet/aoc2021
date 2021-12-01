package utils

import (
	"bufio"
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
