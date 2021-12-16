package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"

	"aoc/2021/utils"
)

type coordinates struct {
	x, y int
}

type graph struct {
	nodes []*node
	start *node
	end   *node
}

type node struct {
	value      int
	neighbours []*node

	dist  int
	index int
}

type queue []*node

func (h queue) Len() int           { return len(h) }
func (h queue) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h queue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *queue) Push(x interface{}) {
	n := len(*h)
	node := x.(*node)
	node.index = n
	*h = append(*h)
}

func (h *queue) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

func (g *graph) findShortestPath() {
	q := make(queue, 0, len(g.nodes))
	prev := make(map[*node]*node)

	for i, node := range g.nodes {
		node.dist = math.MaxInt
		node.index = i
		q = append(q, node)
	}
	g.start.dist = 0
	heap.Init(&q)

	for len(q) != 0 {
		u := heap.Pop(&q).(*node)

		if u == g.end {
			break
		}

		for _, v := range u.neighbours {
			if v.index == -1 {
				continue
			}

			alt := u.dist + v.value
			if alt < v.dist {
				v.dist = alt
				heap.Fix(&q, v.index)

				prev[v] = u
			}
		}
	}
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
	graph := readFile(false)
	graph.findShortestPath()

	return graph.end.dist
}

func solvePart2() int {
	graph := readFile(true)
	graph.findShortestPath()

	return graph.end.dist
}

func expandNum(num, row, col, size int, grid map[coordinates]*node) *node {
	var last *node
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			last = &node{value: (num-1+(i+j))%9 + 1}

			grid[coordinates{col + j*size, row + i*size}] = last
		}
	}

	return last
}

func readFile(expand bool) *graph {
	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(currPath + "/d15/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	grid := make(map[coordinates]*node)
	graph := &graph{nodes: make([]*node, 0, 100)}
	var last *node

	row := 0
	for scanner.Scan() {
		numbers := utils.StringsSliceToIntSlice(strings.Split(scanner.Text(), ""))

		col := 0
		for _, number := range numbers {
			if expand {
				last = expandNum(number, row, col, len(numbers), grid)
			} else {
				last = &node{value: number}
				grid[coordinates{col, row}] = last
			}
			col++
		}
		row++
	}

	for c, num := range grid {
		neighbours := []*node{
			grid[coordinates{x: c.x + 1, y: c.y}],
			grid[coordinates{x: c.x - 1, y: c.y}],
			grid[coordinates{x: c.x, y: c.y + 1}],
			grid[coordinates{x: c.x, y: c.y - 1}],
		}
		for _, neighbour := range neighbours {
			if neighbour == nil {
				continue
			}
			num.neighbours = append(num.neighbours, neighbour)
		}
		graph.nodes = append(graph.nodes, num)
	}
	graph.start = grid[coordinates{0, 0}]
	graph.end = last

	return graph
}
