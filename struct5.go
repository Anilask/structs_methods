package main

import (
	"fmt"
)

// Animal has a Name and an age to reprasent an animal.
type Animal struct {
	Name string
	Age  uint
}

// String makes Animal satisfy the String interface.
func (a Animal) String() string {
	return fmt.Sprintf("%v(%d)", a.Name, a.Age)
}
func main() {
	a := Animal{
		Name: "Lion",
		Age:  3,
	}
	my_str := a.String()
	fmt.Println(my_str)
}
