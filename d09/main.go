package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"

	"aoc/2021/utils"
)

type coordinate struct {
	row, col int
}

type point struct {
	depth       int
	partOfBasin bool
	visited     bool
}

type grid map[coordinate]*point

type basinsHeap []int

func (h basinsHeap) Len() int           { return len(h) }
func (h basinsHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h basinsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *basinsHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *basinsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	utils.RunWithTimeMetricsAndPrintOutput(func() string {
		return fmt.Sprintf("First part: %d", solvePart1())
	})
	utils.RunWithTimeMetricsAndPrintOutput(func() string {
		return fmt.Sprintf("Second part: %d", solvePart2())
	})
}

func solvePart1() int {
	grid := readFile()
	totalRiskLevel := 0
	for c, point := range grid {
		if nextItem(c, coordinate{row: 1, col: 0}, grid).depth > point.depth &&
			nextItem(c, coordinate{row: -1, col: 0}, grid).depth > point.depth &&
			nextItem(c, coordinate{row: 0, col: 1}, grid).depth > point.depth &&
			nextItem(c, coordinate{row: 0, col: -1}, grid).depth > point.depth {
			totalRiskLevel += point.depth + 1
		}
	}

	return totalRiskLevel
}

func solvePart2() int {
	grid := readFile()
	h := &basinsHeap{}

	for c, point := range grid {
		if nextItem(c, coordinate{row: 1, col: 0}, grid).depth > point.depth &&
			nextItem(c, coordinate{row: -1, col: 0}, grid).depth > point.depth &&
			nextItem(c, coordinate{row: 0, col: 1}, grid).depth > point.depth &&
			nextItem(c, coordinate{row: 0, col: -1}, grid).depth > point.depth {
			var sum int
			findBasin(c, grid, &sum)
			heap.Push(h, sum)
		}
	}
	a := heap.Pop(h).(int)
	b := heap.Pop(h).(int)
	c := heap.Pop(h).(int)

	return a * b * c
}

func findBasin(c coordinate, grid grid, sum *int) {
	directions := []coordinate{{row: 1, col: 0}, {row: -1, col: 0}, {row: 0, col: 1}, {row: 0, col: -1}}

	for _, direction := range directions {
		if nextItem := nextItem(c, direction, grid); !nextItem.visited && nextItem.depth < 9 {
			nextItemCoordinates := addCoordinates(c, direction)
			grid[nextItemCoordinates].partOfBasin = true
			grid[nextItemCoordinates].visited = true
			*sum++
			findBasin(nextItemCoordinates, grid, sum)
		}
	}
}

func nextItem(current, direction coordinate, grid grid) *point {
	if point, exists := grid[addCoordinates(current, direction)]; exists {
		return point
	}

	return &point{depth: 10}
}

func addCoordinates(c1, c2 coordinate) coordinate {
	return coordinate{row: c1.row + c2.row, col: c1.col + c2.col}
}

func readFile() grid {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path + "/d09/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	grid := make(grid)
	row := 0
	for scanner.Scan() {
		heights := utils.StringsSliceToIntSlice(strings.Split(scanner.Text(), ""))
		for col, height := range heights {
			grid[coordinate{row: row, col: col}] = &point{depth: height}
		}
		row++
	}

	return grid
}
