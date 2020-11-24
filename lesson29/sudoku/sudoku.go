package sudoku

import (
	"errors"
	"fmt"
)

type Sudoku struct {
	current  [9][9]int8
	original [9][9]int8
}

var (
	ErrFixed           = errors.New("can't change original sudoku digits")
	ErrDupeRow         = errors.New("can't set duplicate digit in row")
	ErrDupeCol         = errors.New("can't set duplicate digit in column")
	ErrDupeSubReg      = errors.New("can't set duplicate digit in subregion")
	ErrCoordOutOfRange = errors.New("row, and/or column is not between 0 and 8, inclusive")
	ErrDigitOutOfRange = errors.New("digit is not between 1 and 9, inclusive")
)

func New(startingDigits [9][9]int8) *Sudoku {
	// The starting digits are fixed in place and may not be overwritten or cleared.
	// Starting digits are saved to original while current is mutable.
	s := Sudoku{
		original: startingDigits,
		current:  startingDigits,
	}

	return &s
}

func (s *Sudoku) PrintBoard() {
	rowCount := len(s.current)

	for i := 0; i < rowCount; i++ {
		fmt.Println(s.current[i])
	}
	fmt.Println("-------------------")
}

func (s *Sudoku) SetDigit(row, col, digit int8) error {
	if !isCoordValid(row, col) {
		return fmt.Errorf("setting %d at row %d, col %d): %w", digit, row, col, ErrCoordOutOfRange)
	}

	if !isDigitValid(digit) {
		return fmt.Errorf("setting %d at row %d, col %d: %w", digit, row, col, ErrDigitOutOfRange)
	}

	if s.isDigitFixed(row, col) {
		return fmt.Errorf("setting %d at row %d, col %d: %w", digit, row, col, ErrFixed)
	}

	if s.isDigitInRow(row, col, digit) {
		return fmt.Errorf("setting %d at row %d, col %d: %w", digit, row, col, ErrDupeRow)
	}

	if s.isDigitInCol(row, col, digit) {
		return fmt.Errorf("setting %d at row %d, col %d: %w", digit, row, col, ErrDupeCol)
	}

	if s.isDigitInSubReg(row, col, digit) {
		return fmt.Errorf("setting %d at row %d, col %d: %w", digit, row, col, ErrDupeSubReg)
	}

	s.current[row][col] = digit

	return nil
}

func (s *Sudoku) ClearDigit(row, col int8) error {
	if !isCoordValid(row, col) {
		return fmt.Errorf("clearing row %d, col %d: %w", row, col, ErrCoordOutOfRange)
	}

	if s.isDigitFixed(row, col) {
		return fmt.Errorf("clearing row %d, col %d: %w", row, col, ErrFixed)
	}

	s.current[row][col] = 0

	return nil
}

func isCoordValid(r, c int8) bool {
	if r < 0 || r > 8 || c < 0 || c > 8 {
		return false
	}

	return true
}

func isDigitValid(d int8) bool {
	if d < 1 || d > 9 {
		return false
	}

	return true
}

// isDigitFixed determines if a digit is fixed by looking at the original board.
func (s *Sudoku) isDigitFixed(r, c int8) bool {
	return s.original[r][c] > 0
}

// Check for a duplicate digit in the same row.
func (s *Sudoku) isDigitInRow(r, c, d int8) bool {
	for c := range s.current[r] {
		if s.current[r][c] == d {
			return true
		}
	}

	return false
}

// Check for duplicate digit in the same column.
func (s *Sudoku) isDigitInCol(r, c, d int8) bool {
	for r := range s.current {
		if s.current[r][c] == d {
			return true
		}
	}

	return false
}

func (s *Sudoku) isDigitInSubReg(r, c, d int8) bool {
	cr, cc := s.findSubRegionCenter(r, c)

	for i := int8(-1); i < 2; i++ {
		for j := int8(-1); i < 2; i++ {
			if s.current[cr+i][cc+j] == d {
				return true
			}
		}
	}

	return false
}

func (s *Sudoku) findSubRegionCenter(r, c int8) (int8, int8) {
	regCenters := [3]int8{1, 4, 7}

	for _, ctrRow := range regCenters {
		for _, ctrCol := range regCenters {
			if abs(ctrRow-r) <= 1 && abs(ctrCol-c) <= 1 {
				return ctrRow, ctrCol
			}
		}
	}

	return -1, -1
}

func abs(n int8) int8 {
	if n < 0 {
		return n * -1
	}

	return n
}
