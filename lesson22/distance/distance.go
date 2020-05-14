/*
Experiment: distance.go
Use the distance method from listing 22.4 to write a program that determines the
distance between each pair of landing sites in table 22.1.

Which two landing sites are the closest?
Which two are farthest apart?

To determine the distance between the following locations, you’ll need to declare
other worlds based on table 22.2:
* Find the distance from London, England (51°30’N 0°08’W) to Paris, France
(48°51’N 2°21’E).
* Find the distance from your city to the capital of your country.
* Find the distance between Mount Sharp (5°4' 48"S, 137°51’E) and Olympus Mons
(18°39’N, 226°12’E) on Mars.

table 22.2

Mercury 2439.7
Venus 6051.8
Earth 6371.0
Mars 3389.5
Jupiter 69911
Saturn 58232
Uranus 25362
Neptune 24622
*/

package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

type location struct {
	lat, long float64
}

type coordinate struct {
	d, m, s float64
	h       rune
}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))

	clong := math.Cos(rad(p1.long - p2.long))

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
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

func main() {
	// mercury := world{2439.7}
	// venus := world{6051.8}
	earth := world{6371.0}
	mars := world{3389.5}
	// jupiter := world{69911}
	// saturn := world{58232}
	// uranus := world{25362}
	// neptune := world{24622}

	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	curiosity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{137, 26, 30.1, 'E'})
	insight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'})

	fmt.Println("Distance (km) between Spirit and Opportunity (farthest)", mars.distance(*opportunity, *spirit))
	fmt.Println("Distance (km) between Spirit and Insight", mars.distance(*insight, *spirit))
	fmt.Println("Distance (km) between Opportunity and Curiosity", mars.distance(*opportunity, *curiosity))
	fmt.Println("Distance (km) between Opportunity and Insight", mars.distance(*opportunity, *insight))
	fmt.Println("Distance (km) between Curiosity and Insight (closest)", mars.distance(*curiosity, *insight))

	london := newLocation(coordinate{51, 30, 0, 'N'}, coordinate{0, 8, 0, 'W'})
	paris := newLocation(coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})
	eugene := newLocation(coordinate{44, 3, 7.449, 'N'}, coordinate{123, 5, 12.313, 'W'})
	washingtonDC := newLocation(coordinate{38, 54, 25.892, 'N'}, coordinate{77, 2, 12.735, 'W'})
	mtSharp := newLocation(coordinate{5, 4, 0, 'S'}, coordinate{137, 51, 0, 'E'})
	olympusMons := newLocation(coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0, 'E'})

	fmt.Println("Distance (km) between London and France", earth.distance(*london, *paris))
	fmt.Println("Distance (km) between Eugene and Washington, D.C.", earth.distance(*eugene, *washingtonDC))
	fmt.Println("Distance (km) between Mount Sharp and Olympus Mons", earth.distance(*mtSharp, *olympusMons))
}
