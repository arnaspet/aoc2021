package main

import (
	"bufio"
	"fmt"
	"os"

	"aoc/2021/utils"
)

func main() {
	utils.RunWithTimeMetricsAndPrintOutput(func() string {
		return fmt.Sprintf("First part: %d", solve(10))
	})
	utils.RunWithTimeMetricsAndPrintOutput(func() string {
		return fmt.Sprintf("Second part: %d", solve(40))
	})
}

func solve(times int) int {
	pairs, replacements := readFile()

	for i := 0; i < times; i++ {
		workMap := make(map[string]int)
		for k, v := range pairs {
			workMap[k] = v
		}

		for pair, times := range pairs {
			if replacement, ok := replacements[pair]; ok {
				workMap[pair] -= times
				if workMap[pair] == 0 {
					delete(workMap, pair)
				}
				workMap[string(pair[0])+replacement] += times
				workMap[replacement+string(pair[1])] += times
			}
		}
		pairs = workMap
		sum := 0
		for _, value := range pairs {
			sum += value
		}
	}

	counts := make(map[uint8]int)
	for pair, value := range pairs {
		counts[pair[1]] += value
	}
	delete(counts, ' ')

	var min, max int
	for _, count := range counts {
		if count > max {
			max = count
		}
		if count < min || min == 0 {
			min = count
		}
	}

	return max - min
}

func solvePart2() int {
	return 0
}

func readFile() (map[string]int, map[string]string) {
	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(currPath + "/d14/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	pairs := make(map[string]int)
	currentPair := []rune{' ', ' '}
	for _, element := range scanner.Text() {
		currentPair[0] = currentPair[1]
		currentPair[1] = element
		pairs[string(currentPair)]++
	}
	scanner.Scan()

	replacements := make(map[string]string)
	for scanner.Scan() {
		var from, to string
		fmt.Sscanf(scanner.Text(), "%s -> %s", &from, &to)
		replacements[from] = to
	}

	return pairs, replacements
}
