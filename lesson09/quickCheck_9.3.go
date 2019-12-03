// Write a program to print each byte (ASCII character) of "shalom", one character per line.
package main

import (
	"fmt"
)

func quickCheck() {
	word := "shalom"
	for i := 0; i <= 5; i++ {
		fmt.Printf("%c\n", word[i])
	}
}
