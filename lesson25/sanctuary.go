/*
Your task is to create a simulation of the first animal sanctuary on Mars. Make a few
types of animals. Each animal should have a name and adhere to the Stringer interface to
return their name.

Every animal should have methods to move and eat. The move method should return a
description of the movement. The eat method should return the name of a random food
that the animal likes.

Implement a day/night cycle and run the simulation for three 24-hour sols (72 hours).
All the animals should sleep from sunset until sunrise. For every hour of the day, pick
an animal at random to perform a random action (move or eat). For every action, print
out a description of what the animal did.

Your implementation should make use of structures and interfaces.
*/

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Every hour is simulated by time in milliseconds
const simHr = 500

type invertebrate struct {
	name       string
	likedFoods []string
	movement   string
}

type bird struct {
	name       string
	likedFoods []string
	movement   string
}

type reptile struct {
	name       string
	likedFoods []string
	movement   string
}

type animal interface {
	String() string
	move() string
	eat() string
}

type sol struct {
	sunrise int
	sunset  int
}

func main() {
	hornet := invertebrate{
		name:       "Rodrick",
		likedFoods: []string{"cricket", "wasp", "rotten fruit", "oak sap"},
		movement:   "buzzes through the air",
	}

	warbler := bird{
		name:       "Del",
		likedFoods: []string{"berry", "seed", "insect"},
		movement:   "steps around on branches high in trees",
	}

	crocodile := reptile{
		name:       "Geho",
		likedFoods: []string{"tilapia", "frog", "clam", "bird", "smaller crododile"},
		movement:   "swims through muddy waters",
	}

	animals := []animal{hornet, warbler, crocodile}

	summerSol := sol{
		sunrise: 5,
		sunset:  20,
	}

	cycles := 3

	// Implement a day/night cycle simulation for three sols (72 hrs).
	for i := 0; i < cycles; i++ {
		fmt.Println("======== Day", i+1, "=======")
		summerSol.simulate(animals)
	}
}

// Simulate a sol (24 hour period)
func (s sol) simulate(animals []animal) {
	dayHrs := s.sunset - s.sunrise
	nightHrs := 24 - dayHrs

	for i := 0; i < dayHrs; i++ {
		randAnim := animals[rand.Intn(len(animals))]
		randAct := randomAction([]string{"eat", "move"})

		switch randAct {
		case "eat":
			fmt.Println(randAnim, "eats", randAnim.eat())

		case "move":
			fmt.Println(randAnim, randAnim.move())
		}

		time.Sleep(simHr * time.Millisecond)
	}

	names := animalNames(animals)
	
	for i := 0; i < nightHrs; i++ {
		fmt.Println(strings.Join(names, ", "), "sleep")
		time.Sleep(simHr * time.Millisecond)
	}
}

func animalNames(animals []animal) []string {
	names := make([]string, len(animals))
	i := 0

	for _, a := range animals {
		names[i] = a.String()
		i++
	}

	return names
}

func randomAction(acts []string) string {
	return acts[rand.Intn(len(acts))]
}

func (inv invertebrate) String() string {
	return inv.name
}

func (brd bird) String() string {
	return brd.name
}

func (rep reptile) String() string {
	return rep.name
}

func (inv invertebrate) move() string {
	return inv.movement
}

func (brd bird) move() string {
	return brd.movement
}

func (rep reptile) move() string {
	return rep.movement
}

func (inv invertebrate) eat() string {
	return randomItem(inv.likedFoods)
}

func (brd bird) eat() string {
	return randomItem(brd.likedFoods)
}

func (rep reptile) eat() string {
	return randomItem(rep.likedFoods)
}

func randomItem(foods []string) string {
	index := rand.Intn(len(foods))
	return foods[index]
}
