package main

import (
	"fmt"
	"math"
	"os"

	"aoc/2021/utils"
)

func main() {
	utils.RunWithTimeMetricsAndPrintOutput(func () string {
		return fmt.Sprintf("First part: %d\n", solvePart1(readFile()))
	})
	utils.RunWithTimeMetricsAndPrintOutput(func () string {
		return fmt.Sprintf("Second part: %d\n", solvePart2(readFile()))
	})

}

func solvePart1(crabs []int) int {
	fuelSums := make([]int, len(crabs))

	for i, destinationCrab := range crabs {
		for _, crab := range crabs {
			fuelSums[i] += int(math.Abs(float64(destinationCrab - crab)))
		}
	}

	min := math.MaxInt
	for _, fuelSum := range fuelSums {
		if fuelSum < min {
			min = fuelSum
		}
	}

	return min
}

func solvePart2(crabs []int) int {
	max := 0
	for _, crab := range crabs {
		if crab > max {
			max = crab
		}
	}
	fuelSums := make([]int, max + 1)

	for i := 0; i <= max; i++ {
		for _, crab := range crabs {
			steps := int(math.Abs(float64(i - crab)))
			fuelSums[i] += ((steps + 1) * steps) / 2
		}
	}

	min := math.MaxInt
	for _, fuelSum := range fuelSums {
		if fuelSum < min {
			min = fuelSum
		}
	}

	return min
}

func readFile() []int {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path + "/d07/input.txt")
	if err != nil {
		panic(err)
	}

	return utils.ReadLineToIntSlice(f)
}
