// Go structs can be nested.
package main 
import"fmt"
 type Address struct{
	 city string 
	 country string
 }
 type User string{
	 name string 
	 age int
	 address Address
 }
 func main(){
	 p:=User{
		 name:"ANil kumar",
		 age:23,
		 address Address{
			 city: "New York",
			 country :"los angle"
		 },
	 }
	 fmt.Println("name:",p.name)
	 fmt.Println("Age:",p.age)
	 fmt.Println("city:",p.address.city)
	 fmt.Println("Country:",p.address.country)

 }