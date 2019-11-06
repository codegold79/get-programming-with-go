package main

import (
	"fmt"
	"math/rand"
)

// Write a guess-the-number program. Make the computer pick random numbers between
// 1–100 until it guesses your number, which you declare at the top of the program. Dis-
// play each guess and whether it was too big or too small.
func main() {
	var number = 33

	for {
		guess := rand.Intn(100) + 1

		if guess == number {
			fmt.Printf("The computer guessed %v correctly.\n", guess)
			break
		} else if guess > number {
			fmt.Printf("The computer guessed %v (too big).\n", guess)
		} else {
			fmt.Printf("The computer guessed %v (too small).\n", guess)
		}
	}
}
