// Cipher the Spanish message “Hola Estación Espacial Internacional” with ROT13. Mod-
// ify listing 9.7 to use the range keyword. Now when you use ROT13 on Spanish text, char-
// acters with accents are preserved.
package main

import (
	"fmt"
)

func international() {
	msg := "Hola Estación Espacial Internacional"

	for _, c := range msg {
		if c >= 'a' && c <= ('z') {
			c += 13

			if c > 'z' {
				c -= 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c += 13
			if c > 'Z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
}
