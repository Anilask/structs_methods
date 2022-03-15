// simple struct
package main

import "fmt"

type User struct {
	name       string
	occupation string
	age        int
	rooll      int
	adress     string
}

func main() {
	u := User{"john Doe", "gardender", 43,1234,"pdpl"}

	fmt.Printf("%s  is %d years old and he is a %s\n", u.name, u.name, u.occupation,u.rooll,u.adress)
}
