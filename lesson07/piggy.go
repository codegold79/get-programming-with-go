// Write a new piggy bank program that uses integers to track the number of cents rather
// than dollars. Randomly place nickels (5¢), dimes (10¢), and quarters (25¢) into an empty
// piggy bank until it contains at least $20.
// Display the running balance of the piggy bank after each deposit in dollars (for exam-
// ple, $1.05).
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var piggy int32
	
	for piggy < 2000 {
		coin := rand.Intn(3)
		switch coin {
		case 0:
			piggy += 5
			fmt.Print("Added a nickel. New balance is ")
		case 1:
			piggy += 10
			fmt.Print("Added a dime. New balance is ")
		case 2:
			piggy += 25
			fmt.Print("Added a quarter. New balance is ")
		}
		fmt.Printf("$%5.2d.\n", piggy/100)
	}
}
