// Experiment: chess.go
// * Extend listing 16.8 to display all the chess pieces at their starting positions using
// the characters kqrbnp for black pieces along the top and uppercase KQRBNP for white
// pieces on the bottom.
// * Write a function that nicely displays the board.
// * Instead of strings, use [8][8]rune to represent the board. Recall that rune literals are
// surrounded with single quotes and can be printed with the %c format verb.
package main

import "fmt"

type board [8][8]rune

const line = "---------------------------------"

func main() {
	board := buildBoard()
	printBoard(board)
}

func buildBoard() board {
	var board board

	// Set up the first row with black pieces
	board[0] = [8]rune{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'}

	// Set up the last row with white pieces
	board[7] = [8]rune{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'}

	// Second row is filled with black pawns
	for col := range board[1] {
		board[1][col] = 'p'
	}

	// Penultimate row is filled with white pawns
	for col := range board[6] {
		board[6][col] = 'P'
	}

	return board
}

func printBoard(b board) {
	for row := range b {
		fmt.Println(line)
		fmt.Print("|")
		for col := range b[row] {
			if b[row][col] == 0 {
				fmt.Print("   |")
			} else {
				fmt.Printf(" %c |", b[row][col])
			}
		}
		fmt.Println()
	}
	fmt.Println(line)
}
