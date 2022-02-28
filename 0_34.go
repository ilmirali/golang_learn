// Variable arguments funcions with slices
package main

import "fmt"

func sum(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}

func f(v ...int) {
	res := 20
	for _, a := range v {
	res -= a
	}
	fmt.Println(res)
}

func main() {
	s := []int{42, 33, 22, 11}
	sum(2, 4, 6)
	sum(s[0], s[1])
	sum(s...)
	
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x[4:])
	
	v := []int{8,5,3}
	f(v...)
	
	p := []int{1, 2, 4, 6, 8}
	p[2] = p[1] 	    //	1 2 2 6 8
	p[3] = p[2] + p[0] //	1 2 2 3 8
	fmt.Println(p[4]%p[3])
	fmt.Println()
	
	var r [5]bool
	fmt.Println(r[3])
}
