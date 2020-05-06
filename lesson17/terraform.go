// Experiment: terraform.go
// Write a program to terraform a slice of strings by prepending each planet with "New ".
// Use your program to terraform Mars, Uranus, and Neptune.
// Your first iteration can use a terraform function, but your final implementation should
// introduce a Planets type with a terraform method, similar to sort.StringSlice.
package main

import (
	"fmt"
	"strings"
)

type Planets []string

func main() {
	var planets Planets
	// The planets set to be terraformed, referred by index of original planets slice.
	listToTerraform := []int{3, 6, 7}

	// First iteration uses terraform function
	planets = []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	terraform(planets, listToTerraform)
	fmt.Println(planets)

	// Second iteration introduces Planet type with terraform method
	planets = Planets{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	fmt.Println(planets)
	planets.terraform(listToTerraform)
	fmt.Println(planets)
}

func terraform(planets []string, listToTerraform []int) {
	for _, p := range listToTerraform {
		planets[p] = "New " + planets[p]
	}
}

func (ps Planets) terraform(listToTerraform []int) {
	for _, p := range listToTerraform {
		ps[p] = "New " + ps[p]
	}
}

func (ps Planets) String() string {
	ss := []string(ps)
	return strings.Join(ss, ", ")
}
