/*
CAPSTONE: SUDOKU RULES

Sudoku is a logic puzzle that takes place on a 9 x 9 grid (see
en.wikipedia.org/wiki/Sudoku). Each square can contain a digit from 1 through
9. The number zero indicates an empty square.

The grid is divided into nine subregions that are 3 x 3 each. When placing a
digit, it must adhere to certain constraints. The digit being placed may not
already appear in any of the following:
	* The horizontal row it’s placed in
	* The vertical column it’s placed in
	* The 3 x 3 subregion it’s placed in

Use a fixed-size (9 x 9) array to hold the Sudoku grid. If a function or method
needs to modify the array, remember that you need to pass the array with a
pointer.

Implement a method to set a digit at a specific location. This method should
return an error if placing the digit breaks one of the rules.

Also implement a method to clear a digit from a square. This method need not
adhere to these constraints, as several squares may be empty (zero).

Sudoku puzzles begin with some digits already set. Write a constructor function
to prepare the Sudoku puzzle, and use a composite literal to specify the
initial values. Here’s an example:

	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

The starting digits are fixed in place and may not be overwritten or cleared.
Modify your program so that it can identify which digits are fixed and which
are penciled in.Add a validation that causes set and clear to return an error
for any of the fixed digits. The digits that are initially zero may be set,
overwritten, and cleared.

You don’t need to write a Sudoku solver for this exercise, but be sure to test
that all the rules are implemented correctly.
*/

package main

import (
	"fmt"

	"github.com/codegold79/get-programming-with-go/lesson29/sudoku"
)

func main() {

	s := sudoku.New([9][9]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
	fmt.Println("New board")
	s.PrintBoard()

	// Solve Sudoku puzzle
	moves := [][3]int8{
		{0, 2, 4},
		{0, 3, 6},
		{0, 5, 8},
		{0, 6, 9},
		{0, 7, 1},
		{0, 8, 2},
		{1, 1, 7},
		{1, 2, 2},
		{1, 6, 3},
		{1, 7, 4},
		{10, 7, 4}, // Incorrect, expect error.
		{1, 8, 8},
		{2, 0, 1},
		{2, 3, 3},
		{2, 4, 4},
		{2, 5, 2},
		{2, 6, 5},
		{2, 8, 7},
		{3, 1, 5},
		{3, 2, 9},
		{3, 2, 8}, // Incorrect, expect error
		{3, 3, 7},
		{3, 5, 1},
		{3, 6, 4},
		{3, 7, 2},
		{4, 1, 2},
		{4, 2, 6},
		{4, 4, 5},
		{4, 6, 7},
		{4, 7, 9},
		{5, 1, 1},
		{5, 2, 3},
		{5, 3, 9},
		{5, 5, 4},
		{5, 6, 8},
		{5, 7, 5},
		{6, 0, 9},
		{6, 2, 1},
		{6, 3, 5},
		{6, 4, 10}, // Incorrect, expect error
		{6, 4, 3},
		{6, 5, 7},
		{6, 8, 4},
		{7, 0, 2},
		{7, 1, 8},
		{7, 2, 7},
		{7, 6, 6},
		{7, 7, 3},
		{7, -7, 3}, // Incorrect, expect error
		{8, 0, 3},
		{8, 1, 4},
		{8, 2, 5},
		{8, 3, 2},
		{8, 5, 6},
		{8, 6, 1},
	}

	for _, m := range moves {
		err := s.SetDigit(m[0], m[1], m[2])
		if err != nil {
			fmt.Printf("\nSetting digit: %v.", err)
		}
	}

	fmt.Println("\n\nSolved board")
	s.PrintBoard()

	// Remove a digit from each subregion
	removes := [][3]int8{
		{1, 2},
		{0, 3},
		{-3, 3}, // Incorrect: expect error
		{2, 6},
		{3, 2},
		{0, 0}, // Incorrect: expect error
		{4, 4},
		{5, 7},
		{6, -1}, // Incorrect: expect error
		{6, 0},
		{8, 3},
		{6, 8},
	}

	for _, r := range removes {
		err := s.ClearDigit(r[0], r[1])
		if err != nil {
			fmt.Printf("\nClearing digit: %v.", err)
		}
	}

	fmt.Println("\n\nPartially cleared board")
	s.PrintBoard()
}
