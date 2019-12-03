// Write a program that converts strings to Booleans:
// * The strings “true”, “yes”, or “1” are true .
// * The strings “false”, “no”, or “0” are false .
// * Display an error message for any other values.
// TIP The switch statement accepts multiple values per case, as covered in lesson 3.
package main

import (
	"fmt"
)

func main() {
	var tF bool
	input := "no"

	switch input {
	case "true", "yes", "1":
		tF = true
	case "false", "no", "0":
		tF = false
	default:
		fmt.Println("Error: invalid values.")
	}

	fmt.Printf("The input value of %s is %t.\n", input, tF)
}
