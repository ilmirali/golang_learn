package main

import (
"fmt"
//"math"
)

var input string

func main(){
x:= 42
y:= 8

res:= x+y
fmt.Println(res)

res = x-y
fmt.Println(res)

res = x*y
fmt.Println(res)

res = x/y
fmt.Println(res)

res = x%y
fmt.Println(res)

s1:= "привет"
s2:= "мир"
fmt.Println(s1 + ", " + s2 + "!")
//fmt.Println(s1 + s2)

fmt.Println("fmt.Println(x==y)",x==y)
fmt.Println("fmt.Println(x!=y)",x!=y)
fmt.Println("fmt.Println(x>y)",x>y)
fmt.Println("fmt.Println(x<y)",x<y)

fmt.Println("(x!=y) && (x<y) THIS IS ", (x!=y) && (x<y))
fmt.Println("(x!=y) || (x<y) THIS IS ", (x!=y) || (x<y))
fmt.Println("!(x!=y) THIS IS", !(x!=y))

//знак амперсанда & перед именем переменной - он используется для возврата адреса переменной
//fmt.Scanln(&input)
//fmt.Println(input)

fmt.Scan(&input)
fmt.Println(input)
}
