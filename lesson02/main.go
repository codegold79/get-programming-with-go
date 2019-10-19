// Write a program to determine how fast a ship would need to travel (in km/h)
// in order to reach Malacandra in 28 days. Assume a distance of 56,000,000 km.
package main

import (
	"fmt"
)

func main() {
	const distance = 56000000
	var hours = 28 * 24
	var spd = distance/hours
	fmt.Println(spd, "km/hr")
}
