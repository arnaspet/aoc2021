package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("First part: %d\n", solve(80, readFile()))
	fmt.Printf("Second part: %d\n", solve(256, readFile()))
}

func solve(days int, fishes map[int]int) int {
	for d := 1; d <= days; d++ {
		for i := 0; i <= 8; i++ {
			fishes[i - 1] = fishes[i]
		}
		fishes[6]  += fishes[-1]
		fishes[8] = fishes[-1]
	}

	total := 0
	delete(fishes, -1)
	for _, count := range fishes {
		total += count
	}

	return total
}

func readFile() map[int]int {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path + "/d06/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	fishes := toMap(strings.Split(scanner.Text(), ","))

	return fishes
}

func toMap(strings []string) map[int]int {
	fishesPool := make(map[int]int)
	for _, number := range strings {
		integer, err := strconv.Atoi(number)
		if err != nil {
			panic("Bad number")
		}
		fishesPool[integer]++
	}
	return fishesPool
}
