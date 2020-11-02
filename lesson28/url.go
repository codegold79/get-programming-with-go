/*
	Experiment: url.go

	In the Go standard library, there’s a function to parse web addresses (see golang.org/
	pkg/net/url/#Parse). Display the error that occurs when url.Parse is used with an invalid
	web address, such as one containing a space: https://a b.com/.

	Use the %#v format verb with Printf to learn more about the error. Then perform a
	*url.Error type assertion to access and print the fields of the underlying structure.
	NOTE A URL, or Uniform Resource Locator, is the address of a page on the World Wide
	Web.
*/

package main

import (
	"fmt"
	"net/url"
)

func main() {
	sites := []string{
		"https://www.golang.org",
		"htts://www.golang.org",
		"htts://www.golang.org",
		"golang.org",
		"https://go lang.org",
	}

	for _, s := range sites {
		u, err := url.Parse(s)
		fmt.Printf("the url: %+v\n", u)

		if err != nil {
			fmt.Printf("\nParse error: %#v", err)
			if urlErr, ok := err.(*url.Error); ok {
				fmt.Printf("\nError Op field: %v", urlErr.Op)
				fmt.Printf("\nError URL field: %v", urlErr.URL)
				fmt.Printf("\nError Err field: %v", urlErr.Err)
			}
		}

		fmt.Println()
	}
}
