// Experiment: words.go
// Write a function to count the frequency of words in a string of text and return a map
// of words with their counts. The function should convert the text to lowercase, and
// punctuation should be trimmed from words. The strings package contains several
// helpful functions for this task, including Fields, ToLower, and Trim .

// Use your function to count the frequency of words in the following passage and then
// display the count for any word that occurs more than once.

// As far as eye could reach he saw nothing but the stems of the great plants about him
// receding in the violet shade, and far overhead the multiple transparency of huge leaves
// filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever
// he felt able he ran again; the ground continued soft and springy, covered with the same
// resilient weed which was the first thing his hands had touched in Malacandra. Once or
// twice a small red creature scuttled across his path, but otherwise there seemed to be no
// life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned
// and alone in a forest of unknown vegetation thousands or millions of miles beyond the
// reach or knowledge of man.
// 		—C.S. Lewis, Out of the Silent Planet, (see mng.bz/V7nO)
package main

import (
	"fmt"
	"strings"
	"unicode"
)

const passage = "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."

const line = "-------------------"

func main() {
	p := strings.ToLower(passage)
	p = strings.TrimFunc(p, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	words := strings.Fields(p)

	wc := make(map[string]int)

	for i := range words {
		wc[words[i]]++
	}

	displayCount(wc)
}

func displayCount(wc map[string]int) {
	fmt.Println("These words appeared more than once:")
	fmt.Println(line)
	fmt.Printf("| %-7s | %5s |\n", "word", "count")
	fmt.Println(line)
	for k, v:=range wc {
		if v > 1 {
			fmt.Printf("| %-7s | %-5d |\n", k, v)
		}
	}
	fmt.Println(line)
}
