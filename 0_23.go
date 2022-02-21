package main
import "fmt"

type Contact struct{
	name string
	age int}

type Car struct{
    color string 
    brand string
    year int
}

type Coord struct{
	x int
	y int
}
type Dog struct{
	name string
	age int
}

func main() {

a := Coord{7, 5}
fmt.Println(a.x-a.y)

x := Contact{"Kate", 42}
y := Contact{name: "Edit", age: 43}
p := &x
p1 := &Contact{"Anna", 15}
j := &Dog{"Max", 8}

fmt.Println("Dog",j.age)

fmt.Println(p.name, p1)
fmt.Println(x,y)
fmt.Println(x.name, x.age)
}	
