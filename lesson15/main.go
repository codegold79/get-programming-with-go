// CAPSTONE: TEMPERATURE TABLES
//
// Write a program that displays temperature conversion tables. The tables
// should use equals signs ( = ) and pipes ( | ) to draw lines, with a
// header section:
// =======================
// | deg C    | deg F    |
// =======================
// | -40.0    | -40.0    |
// | ...      | ...      |
// =======================
//
// The program should draw two tables. The first table has two columns, with
// deg C in the first column and deg F in the second column. Loop from –40 deg C
// through 100 deg C in steps of 5 deg using the temperature conversion methods
// from lesson 13 to fill in both columns.
//
// After completing one table, implement a second table with the columns
// reversed, converting from deg F to deg C.
//
// Drawing lines and padding values is code you can reuse for any data that
// needs to be displayed in a two-column table. Use functions to separate the
// table drawing code from the code that calculates the temperatures for each
// row.
//
// Implement a drawTable function that takes a first-class function as a parameter
// and calls it to get data for each row drawn. Passing a different function
// to drawTable should result in different data being displayed.
package main

import "fmt"

func main() {
	drawTable(cToFRow, "deg C", "deg F")
	drawTable(fToCRow, "deg F", "deg C")
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 1.8) + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) / 1.8
}

func cToFRow(rowFormat string) {
	for t := -40.0; t <= 100; t += 5 {
		fmt.Printf(rowFormat, t, celsiusToFahrenheit(t))
		fmt.Println("\n-----------------------")
	}
}

func fToCRow(rowFormat string) {
	for t := 212.0; t >= -40; t -= 5 {
		fmt.Printf(rowFormat, t, fahrenheitToCelsius(t))
		fmt.Println("\n-----------------------")
	}
}

type rowGeneratorFunc func(string)

func drawTable(rowGenFn rowGeneratorFunc, header1 string, header2 string) {
	fmt.Println("\n=======================")
	fmt.Printf("| %-9s| %-9s|\n", header1, header2)
	fmt.Println("=======================")
	rowGenFn ("| %-9.2f| %-9.2f|")
}
