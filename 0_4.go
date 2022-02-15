package main

import "fmt"

func main(){

var a,b,c int

fmt.Println("Please, enter three numbers across Space key:")
fmt.Scanln(&a, &b, &c)
fmt.Println("And now you can see the product of these numbers:")
fmt.Println(a*b*c)


}
