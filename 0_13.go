package main

import "fmt"

func welcome(name string){
fmt.Println("Hi,", name)
}

func sum(c string, a, b int){
fmt.Println(c)
fmt.Println(a+b)
}

func sum2(x,y int) int {
return x + y
}

func concat(a,b string) string{
return a + b
}

func swap(x,y int) (int, int) {
return y,x
}

func main(){
welcome("David")
sum("Hi!", 4,3)
fmt.Println(sum2(121,100))
fmt.Println(concat("форте","пьяно"))
fmt.Println(swap(1,2))
}	
