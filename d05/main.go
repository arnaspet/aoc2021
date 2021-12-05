package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type coordinate struct {
	x,y int
}

type grid map[coordinate]int

func main() {
	fmt.Printf("First part: %d\n", firstPart(readFile(false)))
	fmt.Printf("Second part: %d\n", firstPart(readFile(true)))
}

func firstPart(grid grid) int {
	return findOverlaps(grid)
}

func findOverlaps(g grid) int {
	var overlaps int
	for _, val := range g {
		if val > 1 {
			overlaps++
		}
	}

	return overlaps
}

func readFile(includeDiagonals bool) grid {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path + "/d05/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	regexp, err := regexp.Compile(`(\d+),(\d+) -> (\d+),(\d+)`)
	if err != nil {
		panic("Bad regexp pattern")
	}

	grid := make(grid)
	for scanner.Scan() {
		numbers := toInt(regexp.FindStringSubmatch(scanner.Text())[1:])
		vector := getMovementVector(numbers)
		if  !includeDiagonals && vector.x != 0 && vector.y != 0 {
			continue
		}
		fillGrid(grid, numbers, vector)
	}

	return grid
}

func fillGrid(g grid, numbers []int, vector coordinate) {
	for {
		g[coordinate{x: numbers[0], y: numbers[1]}]++

		if numbers[0] == numbers[2] && numbers[1] == numbers[3] {
			return
		}

		numbers[0] += vector.x
		numbers[1] += vector.y
	}
}

func getMovementVector(numbers []int) coordinate {
	vector := coordinate{}
	if numbers[3] != numbers[1] {
		diff := numbers[3] - numbers[1]
		vector.y = diff / Abs(diff)
	}
	if numbers[2] != numbers[0] {
		diff := numbers[2] - numbers[0]
		vector.x = diff / Abs(diff)
	}

	return vector
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func toInt(strings []string) []int {
	integers := make([]int, len(strings))
	for i, number := range strings {
		integer, err := strconv.Atoi(number)
		if err != nil {
			panic("Bad number")
		}
		integers[i] = integer
	}
	return integers
}