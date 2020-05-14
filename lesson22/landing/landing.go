/*
Experiment: landing.go
Use the code from listings 22.1, 22.2, and 22.3 to write a program that declares
a location for each location in table 22.1. Print out each of the locations in
decimal degrees.

table 22.1

Spirit
Columbia Memorial Station
	14°34'6.2" S
	175°28'21.5" E

Opportunity
Challenger Memorial Station
	1°56'46.3" S
	354°28'24.2" E

Curiosity
Bradbury Landing
	4°35'22.2" S
	137°26'30.1" E

Insight
Elysium Planitia
	4°30'0.0" N
	135°54'0" E
*/

package main

import (
	"fmt"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

func main() {
	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	fmt.Println("Spirit", *spirit)
	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	fmt.Println("Opportunity", *opportunity)
	curiosity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{137, 26, 30.1, 'E'})
	fmt.Println("Curiosity", *curiosity)
	insight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'})
	fmt.Println("Insight", *insight)
}

// newLocation from latitude, longitude d/m/s coordinates.
func newLocation(lat, long coordinate) *location {
	return &location{lat.decimal(), long.decimal()}
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
