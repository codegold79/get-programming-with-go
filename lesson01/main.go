package main

import (
	"fmt"
)

func main() {
	a()
	b()
	c()
}

func a () {
	fmt.Println("Hello, Frankie")
}

func b() {
	fmt.Println("Hello, world")
	fmt.Println("Hello, 世界")
}

func c() {
	fmt.Println("こんにちは世界")
}