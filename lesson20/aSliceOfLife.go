// A variation of Conway's Game of Life
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func main() {
	A, B := NewUniverse(), NewUniverse()
	Seed(A)

	for i := 0; i < 1000; i++ {
		fmt.Print("\033c")
		Show(A)
		Step(A, B)
		A, B = B, A
		time.Sleep(1 * time.Second / 10)
	}
}

func NewUniverse() Universe {
	var u Universe
	u = make([][]bool, height)

	for i := range u {
		u[i] = make([]bool, width)
	}

	return u
}

func Show(u Universe) {
	for i := range u {
		for j := range u[i] {
			status := " "
			if u[i][j] == true {
				status = "*"
			}
			fmt.Print(status)
		}
		fmt.Println()
	}
}

func Seed(u Universe) {
	// Seed approximately 25% of each row
	var count int = width / 4
	for i := range u {
		for j := 0; j <= count; j++ {
			u[i][rand.Intn(width)] = true
		}
	}
}

// Alive determines the status of a cell, based on seeded/previous state.
// Status of a cell off the board can be examined, by wrapping around.
func (u Universe) Alive(x, y int) bool {
	if x >= width {
		x = width % x
	}
	if x < 0 {
		x += width
	}
	if y >= height {
		y = height % y
	}
	if y < 0 {
		y += height
	}
	return u[y][x]
}

// Neighbor counts number of live adjacent cells.
func (u Universe) Neighbor(x, y int) int {
	live := 0
	for h := -1; h <= 1; h++ {
		for v := -1; v <= 1; v++ {
			// Skip determining if current cell is alive.
			if h == 0 && v == 0 {
				continue
			}

			// Count if adjacent cell is alive.
			if u.Alive(x+h, y+v) {
				live++
			}
		}
	}
	return live
}

// Determine if a cell will be alive next round by examining its neighbors.
func (u Universe) Next(x, y int) bool {
	liveNeighbors := u.Neighbor(x, y)

	switch {
	case liveNeighbors <= 1:
		// Having 1 or no live neighbors means death.
		return false
	case liveNeighbors == 2:
		// Two neighbors means staying alive, if alive.
		return u.Alive(x, y) && true
	case liveNeighbors == 3:
		// Three neighbors means alive, or coming to life from death.
		return true
	case liveNeighbors > 3:
		// Having more than 3 live neighbors means death.
		return false
	}

	return false
}

func Step(a, b Universe) {
	for v := range a {
		for h := range a[v] {
			b[v][h] = a.Next(h, v)
		}
	}
}
