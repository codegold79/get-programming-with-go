/*
Experiment: positionworker.go

Using listing 31.5 as a starting point, change the code so that the delay time
gets a half a second longer with each move.

Listing 31.5

func worker() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	next := time.After(time.Second)
	for {
		select {
			case <-next:
				pos = pos.Add(direction)
				log.Println("current position is ", pos)
				next = time.After(time.Second)
		}
	}
*/

package main

import (
	"image"
	"log"
	"time"
)

func main() {
	go worker()
	time.Sleep(10 * time.Second)
}

func worker() {
	// Include microsecond in log timestamps.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	next := time.After(time.Second)
	var delayFactor time.Duration

	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			log.Println("current position is", pos)
			next = time.After(time.Second + delayFactor*time.Second/2)
			delayFactor++
		}
	}
}
