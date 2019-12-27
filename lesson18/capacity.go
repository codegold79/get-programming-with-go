// Experiment: capacity.go
// Write a program that uses a loop to continuously append an element to a slice. Print out
// the capacity of the slice whenever it changes. Does append always double the capacity
// when the underlying array runs out of room?
package main

import "fmt"

func main() {
	slice := make([]int, 3)
	sCap := cap(slice)

	fmt.Printf("initial slice %v, len: %v, cap: %v\n", slice, len(slice), cap(slice))
	for i := 1; i < 10000; i++ {
		slice = append(slice, i)
		if sCap != cap(slice) {
			fmt.Printf("%d, ", cap(slice))
			sCap = cap(slice)
		}
	}
}
