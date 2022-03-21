package sudoku

import (
	"math/rand"
	"time"
)

// FillField - returns filled the field
func GenerateField(k int) [9][9]int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	field := [9][9]int{}
	ok := false

	for {
		field, ok = fillField()
		if ok {
			break
		}
	}

	// remove elements from field, depending on the coefficient
	for rowIdx := range field {
		for n := ((9 * (k % 100)) / 100); n > 0; n-- {
			s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
			i := r.Intn(len(s))
			field[rowIdx][s[i]] = 0
			s = removeElementFromSlice(s, i)
		}
	}
	return field
}

func fillField() ([9][9]int, bool) {
	res := [9][9]int{}
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	//TODO: the generation can be optimized - generate 3x3 squares by diagonal of the field
	for rowIdx, rowVal := range res {
		for columnIdx := range rowVal {
			s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			for len(s) > 0 {
				i := r1.Intn(len(s))
				if ok := checkGeneratedValue(res, rowIdx, columnIdx, s[i]); ok {
					res[rowIdx][columnIdx] = s[i]
					break
				}
				s = removeElementFromSlice(s, i)
			}
			if len(s) == 0 {
				return res, false
			}
		}
	}
	return res, true
}

func removeElementFromSlice(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
