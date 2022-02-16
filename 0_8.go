package main

import "fmt"

func main(){
	var i int
	
	fmt.Scan(&i)
	switch {
		case i>1000		: fmt.Println("Apple")
		case i>500 && i<=1000	: fmt.Println("Samsung")
		case i<500		: fmt.Println("Nokia с фонариком")
	}
}
