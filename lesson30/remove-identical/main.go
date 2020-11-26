/*
Experiment: remove-identical.go

It’s boring to see the same line repeated over and over again. Write a pipeline
element (a goroutine) that remembers the previous value and only sends the value
to the next stage of the pipeline if it’s different from the one that came before.
To make things a little simpler, you may assume that the first value is never the
empty string.
*/

package main

import "fmt"

func main() {
	msgs := []string{
		"msg1",
		"msg2",
		"msg1",
		"msg3",
		"msg3",
		"msg3",
		"msg4",
		"msg4",
	}

	c0 := make(chan string)
	go sendMsgs(msgs, c0)

	c1 := make(chan string)
	go filterUnique(c0, c1)

	printMsgs(c1)
}

func sendMsgs(msgs []string, downstream chan string) {
	for _, m := range msgs {
		downstream <- m
	}

	close(downstream)
}

// filterUnique ensures only unique values are passed on.
func filterUnique(upstream, downstream chan string) {
	history := make(map[string]struct{})

	for m := range upstream {
		_, ok := history[m]
		if !ok {
			history[m] = struct{}{}
			downstream <- m
		}
	}

	close(downstream)
}

func printMsgs(upstream chan string) {
	for m := range upstream {
		fmt.Println(m)
	}
}
