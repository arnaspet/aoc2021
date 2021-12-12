package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"aoc/2021/utils"
)

type cave struct {
	name  string
	nodes []*cave
}

type caveSystem struct {
	caves map[string]*cave
	paths []path
}
type path []*cave

func (cs caveSystem) getCave(caveName string) *cave {
	if cave, exists := cs.caves[caveName]; exists {
		return cave
	}

	cs.caves[caveName] = &cave{name: caveName}
	return cs.caves[caveName]
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
	caveSystem := readFile()
	caveSystem.findPaths(path{caveSystem.caves["start"]}, visited)
	return len(caveSystem.paths)
}

func solvePart2() int {
	caveSystem := readFile()
	caveSystem.findPaths(path{caveSystem.caves["start"]}, visitedPart2)

	return len(caveSystem.paths)
}

func (cs *caveSystem) findPaths(currentPath path, visitedHandler func(path, *cave) bool) {
	for _, node := range currentPath[len(currentPath)-1].nodes {
		if node.name == "end" {
			finishedPath := make(path, len(currentPath)+1)
			copy(finishedPath, currentPath)
			finishedPath[len(currentPath)] = node
			cs.paths = append(cs.paths, finishedPath)
			continue
		}
		if !visitedHandler(currentPath, node) {
			nextPath := make(path, len(currentPath)+1)
			copy(nextPath, currentPath)
			nextPath[len(currentPath)] = node
			cs.findPaths(nextPath, visitedHandler)
		}
	}
}

func visited(currentPath path, cave *cave) bool {
	if 'A' <= cave.name[0] && cave.name[0] <= 'Z' {
		return false
	}

	for _, node := range currentPath {
		if node == cave {
			return true
		}
	}

	return false
}

func visitedPart2(currentPath path, cave *cave) bool {
	if cave.name == "start" {
		return true
	}
	if 'A' <= cave.name[0] && cave.name[0] <= 'Z' {
		return false
	}

	smallLetters := make(map[string]int)
	visitedTwiceNode := false
	for _, node := range currentPath {
		if 'A' <= node.name[0] && node.name[0] <= 'Z' {
			continue
		}

		smallLetters[node.name]++
		if smallLetters[node.name] == 2 {
			visitedTwiceNode = true
			break
		}
	}
	if !visitedTwiceNode {
		return false
	}

	for _, node := range currentPath {
		if node == cave {
			return true
		}
	}

	return false
}

func readFile() *caveSystem {
	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(currPath + "/d12/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	caveSystem := &caveSystem{caves: make(map[string]*cave), paths: make([]path, 0, 10)}
	for scanner.Scan() {
		relation := strings.Split(scanner.Text(), "-")
		from, to := caveSystem.getCave(relation[0]), caveSystem.getCave(relation[1])

		from.nodes = append(from.nodes, to)
		to.nodes = append(to.nodes, from)
	}

	return caveSystem
}
