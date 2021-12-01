package main

import (
	"fmt"
	"os"

	"aoc/2021/utils"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path + "/d01/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("First part: %d\n", firstPart(f))
	_, _ = f.Seek(0, 0)
	fmt.Printf("Second part: %d\n", secondPart(f))
}

func secondPart(f *os.File) int {
	numbers := utils.ReadFileToIntSlice(f)
	divisibleNumbersLen := len(numbers) - len(numbers) % 3

	increased := -1
	lastScan := 0

	for i, _ := range numbers[:divisibleNumbersLen] {
		deptMeasurement := numbers[i] + numbers[i + 1] + numbers[i + 2]
		if deptMeasurement > lastScan {
			increased++
		}
		lastScan = deptMeasurement
	}

	return increased
}

func firstPart(f *os.File) int {
	increased := -1
	lastScan := 0
	numbers := utils.ReadFileToIntSlice(f)

	for _, number := range numbers {
		if number > lastScan {
			increased++
		}
		lastScan = number
	}

	return increased
}
