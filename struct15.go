// struct pointer  & operator new key word .pointer dediend eith * operator
package main

import "fmt"

type Point struct {
    x int
    y int
}

func main() {

    p := Point{3, 4}

    p_p := &p

    (*p_p).x = 1
    p_p.y = 2

    fmt.Println(p)
}
// Alternatively, we can create a pointer to a struct with the new keyword.


type User struct {
    name       string
    occupation string
    age        int
}

func main() {

    u := new(User)
    u.name = "Richard Roe"
    u.occupation = "driver"
    u.age = 44

    fmt.Printf("%s is %d years old and he is a %s\n", u.name, u.age, u.occupation)
}