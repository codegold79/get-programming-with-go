/* Experiment: landing.go
Write a program that displays the JSON encoding of the three rover landing sites
in listing 21.8. The JSON should include the name of each landing site and use struct
tags as shown in listing 21.10.

To make the output friendlier, make use of the MarshalIndent function from the json
package.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type location struct {
		Name string  `json:"name"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}

	data, err := json.MarshalIndent(&locations, "*", "  ")
	if err != nil {
		fmt.Printf("unable to marshal into json format: %v", err)
	}

	fmt.Println(string(data))
}
