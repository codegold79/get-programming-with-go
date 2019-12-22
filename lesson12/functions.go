// Experiment: functions.go
// Use the Go Playground at play.golang.org to type in listing 12.1 and declare additional
// temperature conversion functions:
// * Reuse the kelvinToCelsius function to convert 233 deg K to Celsius.
// * Write and use a celsiusToFahrenheit temperature conversion function. Hint: the
//   formula for converting to Fahrenheit is: (c * 9.0 / 5.0) + 32.0.
// * Write a kelvinToFahrenheit function and verify that it converts 0 deg K to
//   approximately –459.67° F.
//
// Did you use kelvinToCelsius and celsiusToFahrenheit in your new function or write an
// independent function with a new formula? Both approaches are perfectly valid.

package main

import "fmt"

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "deg K is", celsius, "deg Celsius.")

	kelvin = 233
	celsius = kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "deg K is", celsius, "deg Celsius.")

	kelvin = 0
	fahrenheit := kelvinToFahrenheit(kelvin)
	fmt.Println(kelvin, "deg K is", fahrenheit, "deg Ceclsius.")
}

// kelvinToCelsius converts deg K to deg C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	return 1.8*c + 32
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}
