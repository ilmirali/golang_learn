package main 

import "fmt"

func double(x int) {
	fmt.Println(x*2)
	}

func main(){
	for i := 4; i > 0; i-- {
	defer double(i)
	}


	for i := 4; i > 0; i-- {
	defer fmt.Println(i)
	}

}


