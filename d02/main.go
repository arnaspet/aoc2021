package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"aoc/2021/utils"
)

type submarine struct {
	x, dept, aim int
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path + "/d02/input.txt")
	if err != nil {
		panic(err)
	}

	commands := utils.ReadFileToStringSlice(f)
	firstSub := firstPart(commands)
	fmt.Printf("First part: %d\n\n", firstSub.x*firstSub.dept)

	secondSub := secondPart(commands)
	fmt.Printf("Second part: %d\n\n", secondSub.x*secondSub.dept)
}

func firstPart(commands []string) *submarine {
	sub := &submarine{}
	for _, command := range commands {
		instructions := strings.Split(command, " ")
		increment, err := strconv.Atoi(instructions[1])
		if err != nil {
			panic("bad input")
		}

		switch instructions[0] {
		case "forward":
			sub.x += increment
		case "down":
			sub.dept += increment
		case "up":
			sub.dept -= increment
		}
	}

	return sub
}

func secondPart(commands []string) *submarine {
	sub := &submarine{}
	for _, command := range commands {
		instructions := strings.Split(command, " ")
		increment, err := strconv.Atoi(instructions[1])
		if err != nil {
			panic("bad input")
		}

		switch instructions[0] {
		case "forward":
			sub.x += increment
			sub.dept += sub.aim * increment
		case "down":
			sub.aim += increment
		case "up":
			sub.aim -= increment
		}
	}

	return sub
}
