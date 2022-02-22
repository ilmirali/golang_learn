// methods 
package main
import "fmt"

type Contact struct {
	name string
	age int
}

func (ptr *Contact) increase(val int) {
	ptr.age +=val
}

//func (c *Car) drive() (speed int){
//}

func main() {
	x := Contact{"James", 42}
	x.increase(8)
	fmt.Println(x.age)
	fmt.Println(x.age)
	fmt.Println(x)
}


