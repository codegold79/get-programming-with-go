// Decipher the quote from Julius Caesar:
// L fdph, L vdz, L frqtxhuhg.
// —Julius Caesar
// Your program will need to shift uppercase and lowercase letters by –3. Remember that
// 'a' becomes 'x' , 'b' becomes 'y' , and 'c' becomes 'z' , and likewise for uppercase letters.
package main

import (
	"fmt"
)

func caesar() {
	quote := "L fdph, L vdz, L frqtxhuhg."
	var c rune
	for _, v := range quote {
		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') {
			if v < 'd' || v < 'D' {
				c = v - 3 + 26
			}
			c = v - 3
		} else {
			c = v
		}
		fmt.Printf("%c", c)
	}
	fmt.Println("\n--Julius Caesar")
}
