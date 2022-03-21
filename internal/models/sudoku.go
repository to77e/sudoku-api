package models

type Sudoku struct {
	UUID  string    `json:"uuid"`
	K     int       `json:"k"`
	Field [9][9]int `json:"field"`
}
