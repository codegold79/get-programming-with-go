package main

import (
	"image"
	"sync"
)

type MarsGrid struct {
	mu     sync.Mutex
	bounds image.Rectangle
	cells  [][]cell
}

type message struct {
	rover     string
	pos       image.Point
	lifeSigns int
}

type cell struct {
	occupier  *Occupier
	lifeSigns int
}

type Occupier struct {
	grid *MarsGrid
	pos  image.Point
}

type RoverDriver struct {
	cmdCh         chan cmd
	occupier      *Occupier
	name          string
	marsMsgBuffer chan []message
	fromRover     chan message
}

type cmd int
