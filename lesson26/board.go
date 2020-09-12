/*
	Extension of the turtle program that will show turtle's location.
*/
package main

import (
	"fmt"
)

const (
	minXDefault int = -5
	maxXDefault int = 6
	minYDefault int = -5
	maxYDefault int = 6
)

type board struct {
	x_min int
	y_min int
	x_max int
	y_max int
	plane [][]string
}

func newBoard() *board {
	b := board{
		x_min: minXDefault,
		x_max: maxXDefault,
		y_min: minYDefault,
		y_max: maxYDefault,
	}

	return &b
}

func (b *board) buildPlane(t *turtle) {
	width := b.x_max - b.x_min
	height := b.y_max - b.y_min

	// Board's coordinates do not have negatives, so 0,0 is not center.
	// Shift turtle location to be in the middle of the board.
	turtle_x := t.x + width/2
	turtle_y := height/2 - t.y

	b.plane = make([][]string, width+1)

	for i := 0; i < width; i++ {
		row := make([]string, height+1)
		fmt.Println()
		for j := 0; j < height; j++ {
			row[j] = "[ ]"
			if i == turtle_y && j == turtle_x {
				row[j] = " X "
			}
			if i == width/2 && j == height/2 {
				row[j] = "0,0"
			}

		}
		b.plane[i] = row
	}
}

func (b *board) showPlane(t *turtle) {
	fmt.Print("\033c")
	ht := len(b.plane) - 1
	wd := len(b.plane[0]) - 1

	for i := 0; i < wd; i++ {
		fmt.Println()
		for j := 0; j < ht; j++ {
			fmt.Print(b.plane[i][j])
		}
	}
	fmt.Printf("\nTurtle location: (%v, %v)\n", t.x, t.y)
}
