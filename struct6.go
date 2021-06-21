package main

import (
	"fmt"
)

type Cities struct {
	name     string
	location [2]int
}

func main() {
	// Create empty slice of struct pointer.
	cities := []*Cities{}

	// Create struct and append it to the slice
	ct := new(Cities)
	ct.name = "London"
	ct.location[0] = 5
	ct.location[1] = 10
	cities = append(cities, ct)

	// Create another strut
	ct = new(Cities)
	ct.name = "sydney"
	ct.location[0] = 34
	ct.location[1] = 51
	cities = append(cities, ct)

	for i := range cities {
		c := cities[i]
		fmt.Println("City:", *c)
	}
}
