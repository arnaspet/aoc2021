package utils

import (
	"strconv"
)

func StringsSliceToIntSlice(strings []string) []int {
	integers := make([]int, len(strings))
	for i, number := range strings {
		integers[i] = ParseInt(number)
	}
	return integers
}

func ParseInt(string string) int {
	integer, err := strconv.Atoi(string)
	if err != nil {
		panic("Bad number")
	}

	return integer
}
