package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type number struct {
	row, col int
	marked bool
}

type Board struct {
	numbers map[string]*number
	rows map[int]int
	cols map[int]int
	won bool
}

func NewBoard() *Board {
	return &Board{numbers: make(map[string]*number), rows: make(map[int]int), cols: make(map[int]int)}
}

func main() {
	boards, numbers := readFile()
	
	// fmt.Printf("First part: %d\n",  firstPart(boards, numbers))
	fmt.Printf("Second part: %d\n",  secondPart(boards, numbers))
}

func firstPart(boards []*Board, numbers []string) int {
	for _, calledNumber := range numbers {
		for _, board := range boards {
			if number, ok := board.numbers[calledNumber]; ok {
				number.marked = true
				if board.rows[number.row] += 1; board.rows[number.row] == 5 {
					return bingo(board, calledNumber)
				}
				if board.cols[number.col] += 1; board.cols[number.col] == 5 {
					return bingo(board, calledNumber)
				}
			}
		}
	}

	return 0
}

func secondPart(boards []*Board, numbers []string) int {
	var boardsWon int

	for _, calledNumber := range numbers {
		for _, board := range boards {
			if number, ok := board.numbers[calledNumber]; ok {
				number.marked = true
				if board.rows[number.row] += 1; board.rows[number.row] == 5 && !board.won {
					if !board.won {
						board.won = true
						boardsWon++
					}
				}
				if board.cols[number.col] += 1; board.cols[number.col] == 5 && !board.won {
					board.won = true
					boardsWon++
				}

				if boardsWon == len(boards) {
					return bingo(board, calledNumber)
				}
			}
		}
	}

	return 0
}

func bingo(board *Board, calledNumber string) int {
	sumOfAllUnmarkedNumbers := 0
	for key, number := range board.numbers {
		if !number.marked {
			numberVal, err := strconv.Atoi(key)
			if err != nil {
				panic("Bad input")
			}
			sumOfAllUnmarkedNumbers += numberVal
		}
	}

	calledNumberVal, err := strconv.Atoi(calledNumber)
	if err != nil {
		panic("Bad input")
	}

	return sumOfAllUnmarkedNumbers * calledNumberVal
}

func readFile() ([]*Board, []string)  {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path + "/d04/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	drawnNumbers := strings.Split(scanner.Text(), ",")
	boards := make([]*Board, 0, 100)

	for scanner.Scan() {
		board := NewBoard()

		for row := 0; row < 5; row++ {
			scanner.Scan()
			numbers := strings.Fields(scanner.Text())

			for column, num := range numbers {
				board.numbers[num] = &number{col: column, row: row}
			}
		}
		boards = append(boards, board)
	}

	return boards, drawnNumbers
}
