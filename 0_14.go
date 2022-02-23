
package main 

import "fmt"
var x = 13

func welcome() {
fmt.Println("You are welcome")
}

func change(x int){
x+=1
}

func main(){
defer welcome()
fmt.Println("Hey!")
fmt.Println("Hey!")
fmt.Println("Hey!")
fmt.Println("Hey!")
fmt.Println("Hey!")
fmt.Println("Hey!")
fmt.Println("Hey!")
fmt.Println("Hey!")
fmt.Prinntln(x)
change()
fmt.Prinntln(x)
}
