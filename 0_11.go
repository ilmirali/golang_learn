package main

import "fmt"

func main(){
	sum:=1
	res:=0

	// Loop elements such as starting a counter and changing a counter can be omitted
	for sum<=1000 {
	res += sum
	sum++
	}
	
	fmt.Println(res)

}	
