package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"aoc/2021/utils"
)

type coordinate struct {
	x, y int
}

type grid struct {
	grid map[coordinate]rune
}

func (c coordinate) diff(c2 coordinate) coordinate {
	return coordinate{x: c.x - c2.x, y: c.y - c2.y}
}

func (c coordinate) moreThan(c2 coordinate) bool {
	return c.x >= c2.x && c.y >= c2.y
}

func (g *grid) fold(through coordinate) {
	for c, _ := range g.grid {
		if !c.moreThan(through) {
			continue
		}
		var newCoordinates coordinate
		if through.y > 0 {
			newCoordinates = coordinate{x: c.x, y: through.y - (c.y - through.y)}
		} else {
			newCoordinates = coordinate{x: through.x - (c.x - through.x), y: c.y}
		}

		g.grid[newCoordinates] = '█'
		delete(g.grid, c)
	}
}

func (g *grid) addDot(c coordinate) {
	g.grid[c] = '█'
}

func (g grid) printOut() {
	fmt.Println("---GRID_START---")
	for y := 0; y <= g.maxY(); y++ {
		for x := 0; x <= g.maxX(); x++ {
			c, ok := g.grid[coordinate{x: x, y: y}]
			if !ok {
				c = ' '
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
	fmt.Println("---GRID_END---")
}

func (g grid) maxX() int {
	max := 0
	for c, _ := range g.grid {
		if c.x > max {
			max = c.x
		}
	}

	return max
}

func (g grid) maxY() int {
	max := 0
	for c, _ := range g.grid {
		if c.y > max {
			max = c.y
		}
	}

	return max
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
	grid, folds := readFile()
	grid.fold(folds[0])
	return len(grid.grid)
}

func solvePart2() int {
	grid, folds := readFile()
	for _, fold := range folds {
		grid.fold(fold)
	}
	grid.printOut()
	return 0
}

func readFile() (*grid, []coordinate) {
	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(currPath + "/d13/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	grid := &grid{grid: make(map[coordinate]rune)}
	folds := make([]coordinate, 0, 10)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		coordinates := strings.Split(text, ",")
		grid.addDot(coordinate{x: utils.ParseInt(coordinates[0]), y: utils.ParseInt(coordinates[1])})
	}
	for scanner.Scan() {
		var axis rune
		var value int

		fmt.Sscanf(scanner.Text(), "fold along %c=%d", &axis, &value)
		if axis == 'x' {
			folds = append(folds, coordinate{x: value})
		} else {
			folds = append(folds, coordinate{y: value})
		}
	}

	return grid, folds
}
