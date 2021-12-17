package main

import (
	"bufio"
	"fmt"
	"os"

	"aoc/2021/utils"
)

type grid struct {
	g          map[coordinates]rune
	maxX, maxY int
}

type coordinates struct {
	x, y int
}
type velocity coordinates

func main() {
	utils.RunWithTimeMetricsAndPrintOutput(func() string {
		return fmt.Sprintf("First part: %d", solvePart1())
	})
	utils.RunWithTimeMetricsAndPrintOutput(func() string {
		return fmt.Sprintf("Second part: %d", solvePart2())
	})
}

func solvePart1() int {
	g := readFile()
	maxY := 0
	for x := 1; x <= 100; x++ {
		for y := 0; y <= 1000; y++ {
			if hit, positions := g.shoot(velocity{x, y}); hit {
				max := getMax(positions)
				if max > maxY {
					maxY = max
				}
			}
		}
	}

	return maxY
}

func solvePart2() int {
	g := readFile()
	hitTarget := 0
	for x := 0; x <= 1000; x++ {
		for y := -1000; y <= 5000; y++ {
			if x == 6 && y == 0 {
				fmt.Println()
			}
			if hit, _ := g.shoot(velocity{x, y}); hit {
				hitTarget++
			}
		}
	}

	return hitTarget
}

func getMax(coords []coordinates) int {
	max := 0
	for _, coord := range coords {
		if coord.y > max {
			max = coord.y
		}
	}

	return max
}

func (g *grid) shoot(v velocity) (bool, []coordinates) {
	pc := coordinates{0, 0}
	positions := make([]coordinates, 0, 10)
	for pc.x <= g.maxX && pc.y >= g.maxY {
		pc.x += v.x
		pc.y += v.y
		v.x = approach0(v.x)
		v.y += -1

		positions = append(positions, pc)
		if g.inTarget(pc) {
			return true, positions
		}
	}

	return false, positions
}

func (g *grid) inTarget(pc coordinates) bool {
	return g.g[pc] == 'T'
}

func approach0(num int) int {
	if num > 0 {
		return max(num-1, 0)
	} else {
		return min(num+1, 0)
	}
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func readFile() *grid {
	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(currPath + "/d17/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	grid := &grid{
		g: make(map[coordinates]rune),
	}
	scanner.Scan()
	var xFrom, xTo, yFrom, yTo int
	fmt.Sscanf(scanner.Text(), "target area: x=%d..%d, y=%d..%d", &xFrom, &xTo, &yFrom, &yTo)
	for i := xFrom; i <= xTo; i++ {
		for j := yFrom; j <= yTo; j++ {
			grid.g[coordinates{i, j}] = 'T'
		}
	}
	grid.g[coordinates{0, 0}] = 'S'
	grid.maxX = xTo
	grid.maxY = yFrom

	return grid
}
