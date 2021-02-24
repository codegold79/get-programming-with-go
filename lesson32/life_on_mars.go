package main

import (
	"image"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// receive messages from Mars
	marsToEarth := make(chan []message)
	go receiveMarsMessages(marsToEarth)

	gridMaxPos := image.Point{10, 10}
	startPos := image.Point{0, 0}
	grid := NewMarsGrid(gridMaxPos)
	rovers := make([]*RoverDriver, 5)

	for i, r := range rovers {
		r = NewRoverDriver(grid.Occupy(startPos), "Rover "+strconv.Itoa(i+1), marsToEarth)
		log.Printf("%s created at starting position %v", r.name, r.occupier.pos)
	}

	time.Sleep(10 * time.Second)
}

const (
	right = cmd(0)
	left  = cmd(1)

	// The length of a Mars day.
	dayLength = 10 * time.Second
	// The length of time per day during which
	// messages can be transmitted from a rover to Earth.
	receiveTimePerDay = 2 * time.Second
)

func NewRoverDriver(ocpr *Occupier, name string, marsToEarth chan []message) *RoverDriver {
	r := RoverDriver{
		cmdCh:         make(chan cmd),
		occupier:      ocpr,
		name:          name,
		marsMsgBuffer: marsToEarth,
		fromRover:     make(chan message),
	}

	go r.processMessages()
	go r.drive()

	return &r
}

func (r *RoverDriver) processMessages() {
	var buffered []message

	for {
		select {
		case msg := <-r.fromRover:
			buffered = append(buffered, msg)
		case r.marsMsgBuffer <- buffered:
			buffered = nil
		}
	}
}

func (r *RoverDriver) drive() {
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
			nextMove = time.After(updateInterval)
			newPos := r.occupier.pos.Add(dir)
			if r.occupier.Move(newPos) {
				log.Printf("%s moved to %v", r.name, r.occupier.pos)
				r.reportLifeSigns()
				break
			}

			dir = randomDir(dir)
			log.Printf("%s at %v blocked. New direction set to %v", r.name, newPos, dir)
		}
	}
}

func (r *RoverDriver) reportLifeSigns() {
	r.occupier.grid.mu.Lock()
	defer r.occupier.grid.mu.Unlock()

	lifeSigns := r.occupier.grid.cells[r.occupier.pos.Y][r.occupier.pos.X].lifeSigns
	msg := message{
		rover:     r.name,
		pos:       r.occupier.pos,
		lifeSigns: lifeSigns,
	}

	if lifeSigns >= 900 {
		r.fromRover <- msg
	}
}

func (r *RoverDriver) Left() {
	r.cmdCh <- left
}

func (r *RoverDriver) Right() {
	r.cmdCh <- right
}

func receiveMarsMessages(incoming chan []message) {
	finished := time.After(receiveTimePerDay)

	for {
		time.Sleep(dayLength - receiveTimePerDay)
		for {
			select {
			case <-finished:
			case msgs := <-incoming:
				for _, m := range msgs {
					log.Printf(
						"earth received report of life sign level %d from %s at %v",
						m.lifeSigns,
						m.rover,
						m.pos,
					)
				}
			}
		}
	}
}

func randomDir(oldDir image.Point) image.Point {
	possibleDirs := [3]int{0, -1, 1}

	newDir := image.Point{
		X: possibleDirs[rand.Intn(3)],
		Y: possibleDirs[rand.Intn(3)],
	}

	// Make sure the new direction is different than before.
	for newDir == oldDir {
		newDir.X = possibleDirs[rand.Intn(3)]
		newDir.Y = possibleDirs[rand.Intn(3)]
	}

	return newDir
}

// NewMarsGrid is a rectangular grid with min point at (0, 0) and max point given.
func NewMarsGrid(maxPos image.Point) *MarsGrid {
	// cells store what is occupied in the grid as well as likelihood of life.
	cells := make([][]cell, maxPos.Y)
	for y := range cells {
		cells[y] = make([]cell, maxPos.X)
		for x := range cells[y] {
			cells[y][x].lifeSigns = rand.Intn(1001)
		}
	}

	return &MarsGrid{
		bounds: image.Rectangle{
			Max: maxPos,
		},
		cells: cells,
	}
}

// Occupy occupies a cell at the given point in the grid. It
// returns nil if the point is already occupied or the point is
// outside the grid. Otherwise it returns a value that can be
// used to move to different places on the grid.
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	// Start loop checking the starting position.
	for g.cells[p.Y][p.X].occupier != nil {
		// If starting position is occupied, try a random one.
		p.X = rand.Intn(g.bounds.Max.X)
		p.Y = rand.Intn(g.bounds.Max.Y)
	}

	ocpr := Occupier{g, p}
	g.cells[p.Y][p.X].occupier = &ocpr

	return &ocpr
}

// Move moves the occupier to a different cell in the grid.
// It reports whether the move was successful
// It might fail because it was trying to move outside
// the grid or because the cell it's trying to move into
// is occupied. If it fails, the occupier remains in the same place.
func (ocpr *Occupier) Move(p image.Point) bool {
	ocpr.grid.mu.Lock()
	defer ocpr.grid.mu.Unlock()

	if ocpr.isInBounds(p) && ocpr.isNotOccupied(p) {
		// leave current cell
		ocpr.grid.cells[ocpr.pos.Y][ocpr.pos.X].occupier = nil

		// move to new cell
		ocpr.pos = p
		ocpr.grid.cells[p.Y][p.X].occupier = ocpr
		return true
	}

	return false
}

func (ocpr *Occupier) isInBounds(p image.Point) bool {
	if p.X < ocpr.grid.bounds.Max.X &&
		p.Y < ocpr.grid.bounds.Max.Y &&
		p.X >= ocpr.grid.bounds.Min.X &&
		p.Y >= ocpr.grid.bounds.Min.Y {
		return true
	}

	return false
}

func (ocpr *Occupier) isNotOccupied(p image.Point) bool {
	if ocpr.grid.cells[p.Y][p.X].occupier == nil {
		return true
	}

	return false
}
