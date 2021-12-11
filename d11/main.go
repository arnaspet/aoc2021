package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"aoc/2021/utils"
)

type coordinate struct {
	row, col int
}

type octopus struct {
	power   int
	flashed bool
}

type grid map[coordinate]*octopus

func (g grid) flash(flashingOctopus coordinate) {
	g[flashingOctopus].flashed = true
	adjacentCoordinates := []coordinate{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	for _, adjacentCoordinate := range adjacentCoordinates {
		adjacentOctopusC := addCoordinates(flashingOctopus, adjacentCoordinate)
		adjacentOctopus, ok := g[adjacentOctopusC]
		if !ok {
			continue
		}

		adjacentOctopus.power++
		if adjacentOctopus.power > 9 && !adjacentOctopus.flashed {
			g.flash(adjacentOctopusC)
		}
	}
}

func addCoordinates(c1, c2 coordinate) coordinate {
	return coordinate{row: c1.row + c2.row, col: c1.col + c2.col}
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
	flashSum := 0

	for i := 1; i <= 100; i++ {
		for _, octopus := range grid {
			octopus.power++
		}
		for coordinates, octopus := range grid {
			if octopus.power > 9 && !octopus.flashed {
				grid.flash(coordinates)
			}
		}

		for _, octopus := range grid {
			if octopus.flashed {
				octopus.power = 0
				octopus.flashed = false
				flashSum++
			}
		}
	}

	return flashSum
}

func solvePart2() int {
	grid := readFile()
	step := 1

	for {
		flashCount := 0
		for _, octopus := range grid {
			octopus.power++
		}
		for coordinates, octopus := range grid {
			if octopus.power > 9 && !octopus.flashed {
				grid.flash(coordinates)
			}
		}

		for _, octopus := range grid {
			if octopus.flashed {
				octopus.power = 0
				octopus.flashed = false
				flashCount++
			}
		}
		if len(grid) == flashCount {
			return step
		}
		step++
	}
}

func readFile() grid {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path + "/d11/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	grid := make(grid)
	row := 0
	for scanner.Scan() {
		heights := utils.StringsSliceToIntSlice(strings.Split(scanner.Text(), ""))
		for col, power := range heights {
			grid[coordinate{row: row, col: col}] = &octopus{power: power}
		}
		row++
	}

	return grid
}
