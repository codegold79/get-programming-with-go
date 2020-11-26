/*
Experiment: split-words.go

Sometimes it’s easier to operate on words than on sentences. Write a pipeline
element that takes strings, splits them up into words (you can use the Fields
function from the strings package), and sends all the words, one by one, to the
next pipeline stage.
*/

// take string
// split up into words (strings.Fields)
// send words to next stage

package main

import (
	"fmt"
	"strings"
)

func main() {
	sentences := []string{
		"Yes.",
		"Source: Quanta Magazine by Kevin Hartnett. Four thousand years ago, the Babylonians invented multiplication. Now, mathematicians have perfected it...On March 18, 2019, two researchers described the fastest method ever discovered for multiplying two very large numbers.",
		"Karatsuba’s method made it possible to multiply numbers using only n1.58 single-digit multiplications.",
		"Then in 1971 Arnold Schönhage and Volker Strassen published a method capable of multiplying large numbers in n × log n × log(log n) multiplicative steps, where log n is the logarithm of n.",
		"In 2007 Fürer beat it and the floodgates opened. Over the past decade, mathematicians have found successively faster multiplication algorithms, each of which has inched closer to n × log n, without quite reaching it. Then last month, Harvey and van der Hoeven got there.",
		"Harvey and van der Hoeven’s algorithm proves that multiplication can be done in n × log n steps.",
		"Hardware changes with the times, but best-in-class algorithms are eternal. Regardless of what computers look like in the future, Harvey and van der Hoeven’s algorithm will still be the most efficient way to multiply.",
	}

	c0 := make(chan string)
	go send(sentences, c0)

	c1 := make(chan string)
	go split(c0, c1)

	print(c1)
}

func send(sentences []string, downstream chan string) {
	for _, s := range sentences {
		downstream <- s
	}
	close(downstream)
}

func split(upstream, downstream chan string) {
	for s := range upstream {
		words := strings.Fields(s)
		for _, w := range words {
			downstream <- w
		}
	}

	close(downstream)
}

func print(upstream chan string) {
	for s := range upstream {
		fmt.Println(s)
	}
}
