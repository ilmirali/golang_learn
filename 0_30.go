//slices & make() function
package main
import "fmt"

func main() {
	a := [5]int{0,2,4,6,8}
	fmt.Println(a)	
	b := make([]int,5)
	var 
	s[]int = a[1:3]
	fmt.Println(a)
	s[0] = 8

	//fmt.Println(a, s, a[:3], s[0])
	fmt.Println(a)

	b = append(b,8)
	fmt.Println(b)
}


