//range
package main
import "fmt"

func main() {
	sum := 0
	a := make([]int, 5)
	a[1] = 2
	a[2] = 3
	
	for i,v := range a {
		fmt.Println(i, v)
	}

	for v := range a {
		sum += v
	}

	res := 0
	nums := [3]int{2,4,6}
	for i,v := range nums {
		if i%2 == 0{
			res += v
		}
	}
	
	fmt.Println(sum)
	fmt.Println(0%2==0)
}


