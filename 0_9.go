package main

import "fmt"

func main(){
var res, i,j int
	
	res = 1
	fmt.Scan(&j)
	for i=1; i<=j; i++ {
		res*=i
	}

fmt.Println(res)
}

