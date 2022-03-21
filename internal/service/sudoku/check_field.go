package sudoku

func checkGeneratedValue(field [9][9]int, rowIdx, columnIdx, value int) bool {
	if ok := checkValueInRow(field, rowIdx, value); !ok {
		return false
	}
	if ok := checkValueInColumn(field, columnIdx, value); !ok {
		return false
	}
	if ok := checkValueBySquares(field, rowIdx, columnIdx, value); !ok {
		return false
	}
	return true
}

func checkValueInRow(field [9][9]int, rowIdx, value int) bool {
	for _, v := range field[rowIdx] {
		if v == value {
			return false
		}
	}
	return true
}

func checkValueInColumn(field [9][9]int, columnIdx, value int) bool {
	for _, v := range field {
		if v[columnIdx] == value {
			return false
		}
	}
	return true
}

func checkValueBySquares(field [9][9]int, rowIdx, columnIdx, value int) bool {
	square := [3][3]int{}
	// checking the value for the first square
	if ((rowIdx >= 0) && (rowIdx < 3)) && ((columnIdx >= 0) && (columnIdx < 3)) {
		for rowIdx, rowVal := range field[:3] {
			for columnIdx, columnVal := range rowVal[:3] {
				square[rowIdx][columnIdx] = columnVal
			}
		}
	}
	// checking the value for the second square
	if ((rowIdx >= 0) && (rowIdx < 3)) && ((columnIdx >= 3) && (columnIdx < 6)) {
		for rowIdx, rowVal := range field[:3] {
			for columnIdx, columnVal := range rowVal[3:6] {
				square[rowIdx][columnIdx] = columnVal
			}
		}
	}
	// checking the value for the third square
	if ((rowIdx >= 0) && (rowIdx < 3)) && ((columnIdx >= 6) && (columnIdx < 9)) {
		for rowIdx, rowVal := range field[:3] {
			for columnIdx, columnVal := range rowVal[6:] {
				square[rowIdx][columnIdx] = columnVal
			}
		}
	}
	// checking the value for the fourth square
	if ((rowIdx >= 3) && (rowIdx < 6)) && ((columnIdx >= 0) && (columnIdx < 3)) {
		for rowIdx, rowVal := range field[3:6] {
			for columnIdx, columnVal := range rowVal[:3] {
				square[rowIdx][columnIdx] = columnVal
			}
		}
	}
	// checking the value for the fifth square
	if ((rowIdx >= 3) && (rowIdx < 6)) && ((columnIdx >= 3) && (columnIdx < 6)) {
		for rowIdx, rowVal := range field[3:6] {
			for columnIdx, columnVal := range rowVal[3:6] {
				square[rowIdx][columnIdx] = columnVal
			}
		}
	}
	// checking the value for the sixth square
	if ((rowIdx >= 3) && (rowIdx < 6)) && ((columnIdx >= 6) && (columnIdx < 9)) {
		for rowIdx, rowVal := range field[3:6] {
			for columnIdx, columnVal := range rowVal[6:] {
				square[rowIdx][columnIdx] = columnVal
			}
		}
	}
	// checking the value for the seventh square
	if ((rowIdx >= 6) && (rowIdx < 9)) && ((columnIdx >= 0) && (columnIdx < 3)) {
		for rowIdx, rowVal := range field[6:] {
			for columnIdx, columnVal := range rowVal[:3] {
				square[rowIdx][columnIdx] = columnVal
			}
		}
	}
	// checking the value for the eighth square
	if ((rowIdx >= 6) && (rowIdx < 9)) && ((columnIdx >= 3) && (columnIdx < 6)) {
		for rowIdx, rowVal := range field[6:] {
			for columnIdx, columnVal := range rowVal[3:6] {
				square[rowIdx][columnIdx] = columnVal
			}
		}
	}
	// checking the value for the ninth square
	if ((rowIdx >= 6) && (rowIdx < 9)) && ((columnIdx >= 6) && (columnIdx < 9)) {
		for rowIdx, rowVal := range field[6:] {
			for columnIdx, columnVal := range rowVal[6:] {
				square[rowIdx][columnIdx] = columnVal
			}
		}
	}
	return checkValueInSquare(square, value)
}

func checkValueInSquare(square [3][3]int, value int) bool {
	for rowIdx, rowVal := range square {
		for columnIdx := range rowVal {
			if value == square[rowIdx][columnIdx] {
				return false
			}
		}
	}
	return true
}
