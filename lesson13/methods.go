// Experiment: methods.go
// Write a program with celsius , fahrenheit , and kelvin types and methods to convert from
// any temperature type to any other temperature type.
package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

func main() {
	var c celsius = 100
	var f fahrenheit = -40
	var k kelvin = 0
	fmt.Printf("\n%.2f deg C is %.2f deg F", c, c.fahrenheit())
	fmt.Printf("\n%.2f deg C is %.2f deg K", c, c.kelvin())
	fmt.Printf("\n%.2f deg K is %.2f deg C", k, k.celsius())
	fmt.Printf("\n%.2f deg K is %.2f deg F", k, k.fahrenheit())
	fmt.Printf("\n%.2f deg F is %.2f deg C", f, f.celsius())
	fmt.Printf("\n%.2f deg F is %.2f deg K\n", f, f.kelvin())
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) / 1.8)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*1.8 + 32)
}

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}
