package main

import "fmt"

func main() {
var y string

x := 42
b := &x
*b = 10

//fmt.Println("x", x)
//fmt.Println("b", b, "*b", *b)

a:=  4
fmt.Println(a)
p:=  &a
fmt.Println(a)
fmt.Println(*p)
a += 2
fmt.Println(a)
fmt.Println(*p)
*p = *p-1
fmt.Println(a)
fmt.Println(*p)


fmt.Scan(&y)
fmt.Println(&y)

}
