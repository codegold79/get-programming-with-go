package main

import (
	"image"
	"log"
	"time"
)

type RoverDriver struct {
	cmdCh chan cmd
}

func NewRoverDriver() *RoverDriver {
	r := RoverDriver{
		cmdCh: make(chan cmd),
	}

	go r.drive()

	return &r
}

type cmd int

const (
	right = cmd(0)
	left  = cmd(1)
)

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	dir := image.Point{X: 1, Y: 0}
	updateInterval := time.Millisecond * 250
	nextMove := time.After(updateInterval)

	for {
		select {
		case c := <-r.cmdCh:
			switch c {
			case right:
				dir = image.Point{
					X: dir.Y,
					Y: -dir.X,
				}
			case left:
				dir = image.Point{
					X: -dir.Y,
					Y: dir.X,
				}
			}
			log.Printf("new direction %v", dir)
		case <-nextMove:
			pos = pos.Add(dir)
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func (r *RoverDriver) Left() {
	r.cmdCh <- left
}

func (r *RoverDriver) Right() {
	r.cmdCh <- right
}

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}
