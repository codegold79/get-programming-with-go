/*
Experiment: rover.go

Using the RoverDriver type as a starting point, define Start and Stop methods and
associated commands and make the rover obey them.

// RoverDriver drives a rover around the surface of Mars.
type RoverDriver struct {
	commandc chan command
}
*/

package main

import (
	"image"
	"log"
	"time"
)

type command int

// RoverDriver drives a rover around the surface of Mars.
type RoverDriver struct {
	commandc chan command
}

const (
	stop  = command(0)
	start = command(1)
)

func main() {
	r := NewRoverDriver()
	time.Sleep(time.Second * 1)
	r.Start()
	time.Sleep(time.Second * 1)
	r.Stop()
	time.Sleep(time.Second * 1)
	r.Start()
	time.Sleep(time.Second * 1)
	r.Stop()
	time.Sleep(time.Second * 1)
}

func NewRoverDriver() *RoverDriver {
	r := RoverDriver{
		commandc: make(chan command),
	}

	go r.drive()

	return &r
}

func (r *RoverDriver) Start() {
	r.commandc <- start
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	dir := image.Point{X: 0, Y: 0}
	updateInterval := time.Millisecond * 250
	move := time.After(updateInterval)

	for {
		select {
		case c := <-r.commandc:
			switch c {
			case 0:
				dir = image.Point{X: 0, Y: 0}
				log.Println("Stopped")
			case 1:
				dir = image.Point{X: 1, Y: 0}
				log.Println("Started")
			}
		case <-move:
			pos = pos.Add(dir)
			log.Println("Current position:", pos)
			move = time.After(updateInterval)
		}
	}
}
