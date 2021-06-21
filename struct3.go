// adding method to struct
package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstname string
	lastname  string
	city      string
	gender    string
	age       int
}

func (p Person) hello() string {
	return "Hello , I am " + p.firstname + "" + p.lastname + "," + strconv.Itoa(p.age) + "years old"
}
func main() {
	// initialize person using struct
	p1 := Person{firstname: "Akil", lastname: "kanna", city: "hy", gender: "M", age: 23}
	p2 := Person{"Anilkumar ", "ask", "hyd", "f", 23}
	fmt.Println(p1.hello())
	fmt.Println(p2.hello())

}
