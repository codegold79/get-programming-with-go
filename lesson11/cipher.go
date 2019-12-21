// Experiment: cipher.go
// To send ciphered messages, write a program that ciphers plain text using a keyword:
// plainText := "your message goes here"
// keyword := "GOLANG"
// Bonus: rather than write your plain text message in uppercase letters with no spaces,
// use the strings.Replace and strings.ToUpper functions to remove spaces and uppercase the
// string before you cipher it.
// Once you’ve ciphered a plain text message, check your work by deciphering the
// ciphered text with the same keyword.
// Use the keyword "GOLANG" to cipher a message and post it to the forums for Get Program-
// ming with Go at forums.manning.com/forums/get-programming-with-go.
package main

import (
	"strings"
)

func cipher(plainMsg string, keyword string) string {
	msg := makeWordsUpperCaseWithoutSpaces(plainMsg)
	kw := makeWordsUpperCaseWithoutSpaces(keyword)

	// Repeated keyword
	var kk string

	// Ensure keyword is as long as the message or longer.
	for len(kk) < len(msg) {
		kk += kw
	}

	// Source: https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher#Algebraic_description
	// C_i = E_K(M_i) = (M_i + K_i) mod 26
	// C is the encrypted message, E is the encryption function, K is the keyword,
	// M is the message, and i is the letter index in the word.
	var cipher []byte
	for i := 0; i < len(msg); i++ {
		// Make A=0, B=1, etc.
		m := msg[i] - byte('A')
		k := kk[i] - byte('A')

		c := (m+k)%26 + 'A'
		cipher = append(cipher, c)
	}

	return string(cipher)
}

func makeWordsUpperCaseWithoutSpaces(words string) string {
	w := strings.ToUpper(words)
	w = strings.Replace(w, " ", "", -1)
	return w
}
