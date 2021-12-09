package utils

import (
	"strconv"
)

func StringsSliceToIntSlice(strings []string) []int {
	integers := make([]int, len(strings))
	for i, number := range strings {
		integer, err := strconv.Atoi(number)
		if err != nil {
			panic("Bad number")
		}
		integers[i] = integer
	}
	return integers
}
