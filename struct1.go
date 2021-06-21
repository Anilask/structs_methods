package main
import (
	"fmt"
)
// taking struct
type Rect struct {
	len,wid int
}
func (re Rect) Area() int{
	return re.len*re.wid
}
func main(){
	r:=Rect{10,12}
	fmt.Println("Lengath of the area",r)
	fmt.Println("Area of the Rectangle :",r.Area())
}