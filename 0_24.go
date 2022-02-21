package main

import "fmt"

type Contact struct{
	name string
	age int
}

func (x Contact) welcome() {
	fmt.Println(x.name)
	fmt.Println(x.age)
}

func main() {
	y := Contact{"Dafna", 49}
	//fmt.Println(y)
	x := Contact{"Джеймс", 42}
	x.welcome()
	y.welcome()
}
