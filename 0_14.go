
package main 

import "fmt"

func welcome() {
fmt.Println("You are welcome")
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

}
