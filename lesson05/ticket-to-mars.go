package main

import (
	"fmt"
	"math/rand"
)

// Start by building a prototype that generates 10 random tickets and displays
// them in a tabular format with a nice header...The table should have four columns:
// * The spaceline company providing the service
// * The duration in days for the trip to Mars (one-way)
// * Whether the price covers a return trip
// * The price in millions of dollars
// For each ticket, randomly select one of the following spacelines: Space Adventures,
// SpaceX, or Virgin Galactic.
// Use October 13, 2020 as the departure date for all tickets. Mars will be 62,100,000 km
// away from Earth at the time.
// Randomly choose the speed the ship will travel, from 16 to 30 km/s. This will determine
// the duration for the trip to Mars and also the ticket price. Make faster ships more
// expensive, ranging in price from $36 million to $50 million. Double the price for round trips.
func main() {
	// Distance to Mars in km on Oct 13, 2020
	const distance = 62100000

	// Begin the report with header
	fmt.Println("Spaceline         Days  Trip type   Price")
	fmt.Println("=========================================")

	// Spaceline, if trip is round-trip or one-way
	var company, tripType string

	//  Determine price (millions USD), speed of the ship (km/s), and duration of trip (days)
	var price, speed, duration int

	for i := 0; i < 10; i++ {
		switch num := rand.Intn(3); num {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX          "
		case 2:
			company = "Virgin Galactic "
		}

		// Get spaceship speed which will vary from 16 - 30 km/s
		speed = rand.Intn(14) + 16 // km/s

		// Min - max price = $14M
		price = speed - 16 + 36

		// Get one way duration in days
		duration = distance / (speed * 3600 * 24)

		if num := rand.Intn(2); num == 0 {
			tripType = "One-way   "
		} else {
			tripType = "Round-trip"
			price *= 2
		}

		fmt.Printf("%s   %v   %s   %v\n", company, duration, tripType, price)
	}
	fmt.Println("=========================================")
}
