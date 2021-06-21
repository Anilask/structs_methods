package main

import (
	"fmt"
	"strconv"
)

type person struct {
	firstname string
	lastname  string
	city      string
	gender    string
	age       int
}

// method value receiver
func (p person) hello() string {
	return "Hello, I am" + p.firstname + "" + p.lastname + "" + "," +
		strconv.Itoa(p.age) + "yeras old"
}

// method (pointer receiver) -this modifys the data
func (p *person) hasBirthday() {
	p.age++
}
func main() {
	// initialize person using struct
	p1 := person{firstname: "naveen", lastname: "bottle", city: "hyd", gender: "m", age: 23}
	p2 := person{"neena", "kochi", "kerala", "m", 23}

	fmt.Println(p1)
	fmt.Println(p2)

	p2.hasBirthday() //age +1
	fmt.Println(p2.hello())

	p1.hasBirthday() //age+1
	fmt.Println(p1.hello())

}
