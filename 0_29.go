//array&slices
package main
import "fmt"

func main() {
	var a[5]float32
		
	b := [5]int{0, 2, 4, 6, 8}
	s := b[:]

	//c : =[...]int{0, 2, 4, 6, 8}
	fmt.Println(a,b,s)
	//fmt.Println(c)
	fmt.Println(b[1])
}


