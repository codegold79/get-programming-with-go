// Experiment: terraform.go
// Write a program to terraform a slice of strings by prepending each planet with "New " .
// Use your program to terraform Mars, Uranus, and Neptune.
// Your first iteration can use a terraform function, but your final implementation should
// introduce a Planets type with a terraform method, similar to sort.StringSlice.
package main

import "fmt"

type Planets []string

func main() {
	var planets Planets
	// The planets set to be terraformed, referred by index of original planets slice.
	listToTerraform := []int{3, 6, 7}

	// First iteration uses terraform function
	planets = Planets{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	terraform(planets, listToTerraform)
	fmt.Println(planets)

	// Second iteration introduces Planet type with terraform method
	planets = Planets{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	fmt.Println(planets)
	planets.terraform(listToTerraform)
	fmt.Println(planets)
}

func terraform(planets Planets, listToTerraform []int) {
	for _, p := range listToTerraform {
		planets[p] = "New " + planets[p]
	}
}

func (planets Planets) terraform(listToTerraform []int) {
	for _, p := range listToTerraform {
		planets[p] = "New " + planets[p]
	}
}
