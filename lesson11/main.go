package main

import (
	"fmt"
)

func main() {
	cipherTxt := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	fmt.Println("Decoded using strings package and no range keyword:")
	fmt.Println(decipherWithStringsNoRange(cipherTxt, keyword) + "\n")
	fmt.Println("Decoded using range but no other package except fmt:")
	fmt.Println(decipherWithRangeNoStringsPkg(cipherTxt, keyword) + "\n")
	fmt.Println("Decoded without use of if statements, but modulo:")
	fmt.Println(decipherUseMod(cipherTxt, keyword))

	plainTxt := "your message goes here"
	fmt.Println("\nCipher of the plain text message", plainTxt, "using", keyword, "is: ")
	cipherTxt = (cipher(plainTxt, keyword))
	fmt.Println(cipherTxt)
	fmt.Println("\nDeciphered message is:")
	fmt.Println(decipherUseMod(cipherTxt, keyword))
}
