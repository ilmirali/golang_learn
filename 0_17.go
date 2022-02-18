package main 

import "fmt"

func double_m(a,b int) int {
	sum:=0
	for a<=b {
		sum += a*a
		fmt.Println(a*a)
		a++
	}
	return sum
}

func main(){
fmt.Println(double_m(2,6))
}


