/*
	Write a program with a turtle that can move up, down, left, or right. The turtle should
	store an (x, y) location where positive values go down and to the right. Use methods to
	increment/decrement the appropriate variable. A main function should exercise the
	methods you’ve written and print the final location.

	TIP: Method receivers will need to use pointers in order to manipulate the x and y values.
*/
package main

import (
	"fmt"
	"time"
)

type turtle struct {
	x int
	y int
}

func main() {
	fred := &turtle{}
	board := newBoard()
	board.buildPlane(fred)
	board.showPlane(fred)
	fmt.Println("Build plane")
	fmt.Println("Next: Move turtle right 2")
	waitTime := 3 * time.Second

	fred.move("x", 2)
	board.buildPlane(fred)
	board.showPlane(fred)
	fmt.Println("Next: Move turtle up 4")
	time.Sleep(waitTime)

	fred.move("y", 4)
	board.buildPlane(fred)
	board.showPlane(fred)
	fmt.Println("Move turtle left 5")
	time.Sleep(waitTime)

	fred.move("x", -5)
	board.buildPlane(fred)
	board.showPlane(fred)
	fmt.Println("Move turtle down 1")
	time.Sleep(waitTime)

	fred.move("y", -1)
	board.buildPlane(fred)
	board.showPlane(fred)
	fmt.Println("Move turtle right 2")
	time.Sleep(waitTime)

	fred.move("x", 2)
	board.buildPlane(fred)
	board.showPlane(fred)
	fmt.Println("Move turtle down 6")
	time.Sleep(waitTime)

	fred.move("y", -6)
	board.buildPlane(fred)
	board.showPlane(fred)
	fmt.Println("Move turtle right 3")
	time.Sleep(waitTime)

	fred.move("x", 3)
	board.buildPlane(fred)
	board.showPlane(fred)

	fmt.Println()
}

func (t *turtle) move(axis string, delta int) {
	switch axis {
	case "x":
		t.x += delta
	case "y":
		t.y += delta
	}
}
