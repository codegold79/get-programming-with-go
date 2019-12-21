// Experiment: decipher.go
// * Write a program to decipher the ciphered text
// cipherText := "CSOITEUIWUIZNSROCNKFD"
// keyword := "GOLANG"
// To keep it simple, all characters are uppercase English letters for both the text and keyword.
// * The strings.Repeat function may come in handy. Give it a try, but also complete
// this exercise without importing any packages other than fmt to print the deciphered message.
// * Try this exercise using range in a loop and again without it. Remember that the
// range keyword splits a string into runes, whereas an index like keyword[0] results in
// a byte.
// * You can only perform operations on values of the same type, but you can convert one
// type to the other (string, byte, rune).
// * To wrap around at the edges of the alphabet, the Caesar cipher exercise made use of
// a comparison. Solve this exercise without any if statements by using modulus.
// * If you recall, modulus gives the remainder of dividing two numbers. For example, 27 % 26
// is 1, keeping numbers within the 0–25 range. Be careful with negative numbers, though, as
// -3 % 26 is still -3.
package main

import (
	"strings"
)

func decipherWithStringsNoRange(cipherTxt string, keyword string) string {
	repeatCount := int(strings.Count(cipherTxt, "") / strings.Count(keyword, ""))
	repeatedKeyword := strings.Repeat(keyword, repeatCount+1)
	var decipheredTxt []byte

	for i := 0; i < len(cipherTxt); i++ {
		if cipherTxt[i] > repeatedKeyword[i] {
			c := cipherTxt[i] - repeatedKeyword[i] + 'A'
			// fmt.Printf("%d: %c %c %c %[2]d %[3]d %[4]d\n\n", i, cipherTxt[i], repeatedKeyword[i], c)
			decipheredTxt = append(decipheredTxt, c)
		} else if cipherTxt[i] < repeatedKeyword[i] {
			c := 'Z' - repeatedKeyword[i] + cipherTxt[i] + 1
			// fmt.Printf("%d: %c %c %c %[2]d %[3]d %[4]d\n\n", i, cipherTxt[i], repeatedKeyword[i], c)
			decipheredTxt = append(decipheredTxt, c)
		} else {
			c := byte('A')
			// fmt.Printf("%d: %c %c %c %[2]d %[3]d %[4]d\n", i, cipherTxt[i], repeatedKeyword[i], c)
			decipheredTxt = append(decipheredTxt, c)
		}
	}
	return string(decipheredTxt)
}

func decipherWithRangeNoStringsPkg(cipherTxt string, keyword string) string {
	var i int
	var decoded string

	for _, c := range cipherTxt {
		// If keyword letter counter is at end of keyword, restart at beginning.
		if i == len(keyword) {
			i = 0
		}

		switch {
		case c > rune(keyword[i]):
			decoded += string(c - rune(keyword[i]) + 'A')
		case c < rune(keyword[i]):
			decoded += string('Z' - rune(keyword[i]) + c + 1)
		default:
			decoded += string('A')
		}

		i++
	}

	return decoded
}

func decipherUseMod(cipherTxt string, keyword string) string {
	// Source: https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher#Algebraic_description
	// Given that A=0, B=1, etc.
	// Decryption is done with the equation below:
	// M_i = D_K(C_i) = (C_i - K_i) mod 26
	// M is the decrypted message, D is the formula, K is the keyword,
	// C is the encrypted message, i is the letter index in the associated word.

	var repeatedKW string

	// Keyword should be repeated such that it is longer than the encrypted message
	for len(cipherTxt) > len(repeatedKW) {
		repeatedKW += keyword
	}

	var message []byte

	for i, ct := range cipherTxt {
		// Set the letters such that A=0, B=2, etc.
		c := byte(ct) - 'A'
		k := repeatedKW[i] - 'A'

		// We don't want any negative numbers, so increase the difference
		// between c and k by 26, which does not change modulo result.
		m := (c - k + 26) % 26

		// Return letters back to UTF-8 value before adding back to slice.
		message = append(message, byte(m+65))
	}
	return string(message)
}
