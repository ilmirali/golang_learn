package main 

import "fmt"

func one_or_two (a,b int, s string) (int, string) {
	switch {
	case s == "one": return a, s
	case s == "two": return b, s
	default: return 0, s
	} 
}

func main(){
fmt.Println(one_or_two(1,2,"two"))
}


