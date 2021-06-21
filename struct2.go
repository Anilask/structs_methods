package main
import (
	"fmt"
)
type Person struct {
	firstname string
	lastname string
	city string
	gender string
	age int 
}
func main(){
	// initialize Person using strut 
	p1:=Person{firstname: "Anil",lastname: "kumar",city: "hyd",gender: "male",age: 23}
	p2:=Person{"Ruchitha","sura","hyd","Female",22}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println (p2.firstname,p2.lastname)
	p2.age--
	fmt.Println(p2)
}