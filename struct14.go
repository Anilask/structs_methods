// Go structs can be nested.

package main
import "fmt"
type info func(string,string,int)string
type User struct{
	name string
	occupation string
	age int
	info info
}
func main(){
	u:=User{
		name: "Anil Ask",
		occupation: "B.tech",
		age: 34,
		
		info:func(name string, occupation string,age int)string {
			return fmt.Sprintf("%s is %d years old and he is a %s\n",name,age,occupation)
		},
	}
		fmt.Printf(u.info(u.name,u.occupation,u.age))
}