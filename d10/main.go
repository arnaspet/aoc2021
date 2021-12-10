package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/thoas/go-funk"

	"aoc/2021/utils"
)

// 40 91 123 60
var availableOpeningBraces = []rune{'(', '[', '{', '<'}

// 41 93 125 62
var availableClosingBraces = []rune{')', ']', '}', '>'}

var parenthesisPairs = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',

	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var syntaxErrorScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var missingErrorScores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
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
	lines := readFile()
	syntaxErrors := make([]rune, 0, 10)

	for _, line := range lines {
		openingBraces := make([]rune, 0, 10)

		for _, character := range line {
			if funk.Contains(availableOpeningBraces, character) {
				openingBraces = append(openingBraces, character)
			}
			if funk.Contains(availableClosingBraces, character) {
				if openingBraces[len(openingBraces)-1] == parenthesisPairs[character] {
					openingBraces = openingBraces[:len(openingBraces)-1]
				} else {
					syntaxErrors = append(syntaxErrors, character)
					break
				}
			}
		}
	}

	sum := 0
	for _, syntaxError := range syntaxErrors {
		sum += syntaxErrorScores[syntaxError]
	}

	return sum
}

func solvePart2() int {
	lines := readFile()
	repairScores := make([]int, 0, 10)

	for _, line := range lines {
		openingBraces := make([]rune, 0, 10)
		corrupted := false

		for _, character := range line {
			if funk.Contains(availableOpeningBraces, character) {
				openingBraces = append(openingBraces, character)
			}
			if funk.Contains(availableClosingBraces, character) {
				if openingBraces[len(openingBraces)-1] == parenthesisPairs[character] {
					openingBraces = openingBraces[:len(openingBraces)-1]
				} else {
					corrupted = true
					break
				}
			}
		}
		if len(openingBraces) != 0 && !corrupted {
			// needs repairing
			repairScores = append(repairScores, getRepairScore(openingBraces))
		}
	}

	sort.Ints(repairScores)

	return repairScores[len(repairScores)/2]
}

func getRepairScore(braces []rune) int {
	score := 0

	for i := len(braces) - 1; i > -1; i-- {
		missingCharacter := parenthesisPairs[braces[i]]
		score = score*5 + missingErrorScores[missingCharacter]
	}

	return score
}

func readFile() []string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path + "/d10/input.txt")
	if err != nil {
		panic(err)
	}
	return utils.ReadFileToStringSlice(f)
}
