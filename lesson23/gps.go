/* Experiment: gps.go
Write a program with a gps structure for a Global Positioning System (GPS). This struct
should be composed of a current location, destination location, and a world.

Implement a description method for the location type that returns a string containing the
name, latitude, and longitude. The world type should implement a distance method
using the math from lesson 22.

Attach two methods to the gps type. First, attach a distance method that finds the distance
between the current and destination locations. Then implement a message method that
returns a string describing how many kilometers remain to the destination.

As a final step, create a rover structure that embeds the gps and write a main function to
test everything out. Initialize a GPS for Mars with a current location of Bradbury Land-
ing (-4.5895, 137.4417) and a destination of Elysium Planitia (4.5, 135.9). Then create a
curiosity rover and print out its message (which forwards to the gps ).
*/

package main

import (
	"fmt"
	"math"
)

type location struct {
	name string
	lat  float64
	long float64
}

type world struct {
	radius float64
}

type gps struct {
	current     location
	destination location
	planet      world
}

type rover struct {
	gps
}

func main() {
	mars := world{3389.5}

	g := gps{
		current:     location{"Bradbury Landing", -4.5895, 137.4417},
		destination: location{"Elysium Planitia", 4.5, 135.9},
		planet:      mars,
	}

	curiosity := rover{g}

	fmt.Println("Current", curiosity.current.description())
	fmt.Println("Destination", curiosity.destination.description())
	fmt.Println(curiosity.message())
}

func (loc location) description() string {
	return fmt.Sprintf(
		"location name: %s, coordinates: (%f, %f).",
		loc.name, loc.lat,
		loc.long,
	)
}

// distance between current and destination
func (g gps) distance() float64 {
	return g.planet.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("%f km remain to destination.", g.distance())
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))

	clong := math.Cos(rad(p1.long - p2.long))

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}
