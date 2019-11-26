// Canis Major Dwarf is the closest known galaxy to Earth at 236,000,000,000,000,000 km
// from our Sun (though some dispute that it is a galaxy). Use constants to convert this dis-
// tance to light years.
package main

import (
	"fmt"
)

func main() {
	const distToCanis = 236e15 // km
	const lightSpd = 299792    // km/s
	const secsPerYear = 3600 * 24 * 365.25

	lightYearsToCanis := distToCanis / lightSpd / secsPerYear
	fmt.Printf("Canis Major Dwarf galaxy is %d light years away.\n", int32(lightYearsToCanis))
}
