// Save some money to buy a gift for your friend. Write a program that randomly places
// nickels ($0.05), dimes ($0.10), and quarters ($0.25) into an empty piggy bank until it con-
// tains at least $20.00. Display the running balance of the piggy bank after each deposit,
// formatting it with an appropriate width and precision.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nickel := 0.05
	dime := 0.1
	quarter := 0.25

	var piggy float64

	for piggy < 20 {
		coin := rand.Intn(3)
		switch coin {
		case 0:
			piggy += nickel
			fmt.Print("Added a nickel. New balance is ")
		case 1:
			piggy += dime
			fmt.Print("Added a dime. New balance is ")
		case 2:
			piggy += quarter
			fmt.Print("Added a quarter. New balance is ")
		}
		fmt.Printf("$%5.2f.\n", piggy)
	}
}
