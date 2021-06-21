// the structs fields are accessed with the dot  operator.
package main

import "fmt"

type User struct {
	name       string
	occupation string
	age        int
}

func main() {
	u := User{}
	u.name = "Anil"
	u.occupation = "B.tech"
	u.age = 23
	fmt.Printf("%s is %d years old and he is %s\n", u.name, u.age, u.occupation)
}
