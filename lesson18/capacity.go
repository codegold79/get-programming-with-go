// Experiment: capacity.go
// Write a program that uses a loop to continuously append an element to a slice. Print out
// the capacity of the slice whenever it changes. Does append always double the capacity
// when the underlying array runs out of room?
package main

import "fmt"

func main() {
	slice := make([]int, 1)
	fmt.Printf("initial slice %v, len: %v, cap: %v\n", slice, len(slice), cap(slice))
	for i := 1; i < 11; i++ {
		slice = append(slice, i)
		fmt.Printf("%v: %v, len: %v, cap: %v\n", i, slice, len(slice), cap(slice))
	}
	fmt.Println("It does seem that append always doubles the capacity when the underlying array runs out of room.")
}
