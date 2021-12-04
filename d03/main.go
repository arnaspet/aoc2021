package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"

	"aoc/2021/utils"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Open(path + "/d03/input.txt")
	if err != nil {
		panic(err)
	}

	bits := utils.ReadFileToStringSlice(f)
	fmt.Printf("First part: %d\n", firstPart(bits))
	f.Seek(0, 0)
	fmt.Printf("Second part: %d\n", secondPart(f))
}

func firstPart(bitLines []string) uint {
	counts := make(map[int]map[rune]int)
	dataSize := len(bitLines[0])
	for i :=0; i < dataSize; i++ {
		counts[i] = make(map[rune]int)
	}

	for _, line := range bitLines {
		for j, bit := range line {
			counts[j][bit]++
		}
	}

	var gb, eb strings.Builder
	for i := 0; i < dataSize; i++ {
		if counts[i]['0'] > counts[i]['1'] {
			gb.WriteRune('0')
			eb.WriteRune('1')
		} else {
			gb.WriteRune('1')
			eb.WriteRune('0')
		}
	}
	gamma, err := strconv.ParseInt(gb.String(), 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(eb.String(), 2, 64)
	if err != nil {
		panic(err)
	}

	return uint(gamma) * uint(epsilon)
}

func secondPart(f *os.File) uint {
	listForOxygen := utils.ReadFileToStringList(f)
	f.Seek(0, 0)
	listForCO := utils.ReadFileToStringList(f)

	dataSize := len(listForOxygen.Front().Value.(string))


	for column := 0; column < dataSize && listForOxygen.Len() > 1; column++ {
		elements := make(map[rune][]*list.Element)

		for e := listForOxygen.Front(); e != nil; e = e.Next() {
			bit := e.Value.(string)[column]
			elements[rune(bit)] = append(elements[rune(bit)], e)
		}

		if len(elements['0']) > len(elements['1']) {
			clearLines(elements['1'], listForOxygen)
		} else {
			clearLines(elements['0'], listForOxygen)
		}
	}
	oxygenGeneratorRating, _ := strconv.ParseInt(listForOxygen.Front().Value.(string), 2, 64)

	for column := 0; column < dataSize && listForCO.Len() > 1; column++ {
		elements := make(map[rune][]*list.Element)

		for e := listForCO.Front(); e != nil; e = e.Next() {
			bit := e.Value.(string)[column]
			elements[rune(bit)] = append(elements[rune(bit)], e)
		}

		if len(elements['0']) > len(elements['1']) {
			clearLines(elements['0'], listForCO)
		} else {
			clearLines(elements['1'], listForCO)
		}
	}
	coGeneratorRating, _ := strconv.ParseInt(listForCO.Front().Value.(string), 2, 64)


	return uint(oxygenGeneratorRating) * uint(coGeneratorRating)
}

func clearLines(linesToBeCleared []*list.Element, bitLines *list.List) {
	for _, element := range linesToBeCleared {
		bitLines.Remove(element)
	}
}
