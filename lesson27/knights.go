/*
	Experiment: knights.go
	A knight blocks Arthur’s path. Our hero is empty-handed, represented by a nil value for
	leftHand *item. Implement a character struct with methods such as pickup(i *item) and
	give(to *character). Then use what you’ve learned in this lesson to write a script that has
	Arthur pick up an item and give it to the knight, displaying an appropriate description
	for each action.
*/

package main

import (
	"fmt"
)

type character struct {
	name     string
	leftHand *item
}

type item struct {
	name string
	kind string
}

func main() {
	knight := &character{name: "Scott"}
	arthur := &character{name: "Arthur"}

	fmt.Printf(
		"%s, a knight, stands in the middle of a path.\n", knight.name)
	knight.wields()

	fmt.Printf(
		"%s is walking on the path, but is blocked by %s.\n", arthur.name, knight.name)
	arthur.wields()

	weapon := &item{
		kind: "sword",
		name: "Nexcalibre",
	}

	fmt.Printf(
		"%s sees a sword hiding in the tall grass growing along the sides of the well-trodden path. ",
		arthur.name,
	)

	fmt.Printf(
		"%s ducks into the grass, disappearing momentarily. ",
		arthur.name,
	)

	arthur.pickup(weapon)
	arthur.give(weapon, knight)

	arthur.wields()
	knight.wields()
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}

	fmt.Printf("%s picks up a %v.\n", c.name, i.kind)
	c.leftHand = i
}

func (c1 *character) give(i *item, c2 *character) {
	if c1 == nil || c2 == nil || i == nil {
		return
	}

	if c1.leftHand == nil || c1.leftHand != i {
		fmt.Printf("Sorry, %s does not hold the item to give.\n", c1.name)
		return
	}

	if c2.leftHand != nil {
		fmt.Printf("%s cannot carry another item.\n", c2.name)
		return
	}

	fmt.Printf("%s hands a %v to %s.\n", c1.name, c1.leftHand.kind, c2.name)

	c1.leftHand = nil
	c2.leftHand = i
}

func (c *character) wields() {
	if c == nil {
		return
	}

	if c.leftHand == nil {
		fmt.Printf("%s wields nothing in zir left hand.\n", c.name)
		return
	}

	fmt.Printf("%s wields a %s in zir left hand.\n", c.name, c.leftHand.kind)
}
