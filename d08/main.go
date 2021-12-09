package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"

	"aoc/2021/utils"
)

func main() {
	utils.RunWithTimeMetricsAndPrintOutput(func() string {
		return fmt.Sprintf("First part: %d", solvePart1())
	})
	utils.RunWithTimeMetricsAndPrintOutput(func() string {
		return fmt.Sprintf("Second part: %d", solvePart2())
	})
	solvePart2()
}

func solvePart1() int {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path + "/d08/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " | ")
		for _, digit := range strings.Split(input[1], " ") {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				sum++
			}
		}
	}

	return sum
}

func solvePart2() int {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path + "/d08/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " | ")
		sum += calcSum(strings.Split(input[0], " "), strings.Split(input[1], " "))
	}
	return sum
}

func calcSum(digits []string, sumDigits []string) int {
	parsedDigits := make(map[string][]string)

	i := 0
	for len(parsedDigits) != 10 {
		letters := strings.Split(digits[i], "")
		sort.Strings(letters)
		switch {
		case len(letters) == 2:
			parsedDigits["1"] = letters
		case len(letters) == 3:
			parsedDigits["7"] = letters
		case len(letters) == 4:
			parsedDigits["4"] = letters
		case len(letters) == 7:
			parsedDigits["8"] = letters
		case len(letters) == 6 && len(parsedDigits) >= 4: // 0 6 9
			if intersection := funk.IntersectString(parsedDigits["7"], letters); len(intersection) == 2 {
				parsedDigits["6"] = letters
				break
			}
			if intersection := funk.IntersectString(parsedDigits["4"], letters); len(intersection) == 3 {
				parsedDigits["0"] = letters
			} else {
				parsedDigits["9"] = letters
			}
		case len(letters) == 5 && len(parsedDigits) >= 7: // 2 3 5
			if intersection := funk.IntersectString(letters, parsedDigits["7"]); len(intersection) == 3 {
				parsedDigits["3"] = letters
				break
			}
			if intersection := funk.IntersectString(letters, parsedDigits["6"]); len(intersection) == 5 {
				parsedDigits["5"] = letters
			} else {
				parsedDigits["2"] = letters
			}

		}
		i = (i + 1) % len(digits)
	}

	var sum string
	for _, sumDigit := range sumDigits {
		letters := strings.Split(sumDigit, "")
		sort.Strings(letters)
		for number, parsedDigit := range parsedDigits {
			if Equal(letters, parsedDigit) {
				sum += number
			}
		}
	}
	sumNumber, _ := strconv.Atoi(sum)
	return sumNumber
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
