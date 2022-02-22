// methods with pointers (pointers need to change structure data in the method)
package main
import "fmt"

type Test struct {
	a int
	b int
}

func (x Test) do() int {
	return x.a -x.b
}

func main() {
	t := Test{11, 2} 
	fmt.Println(t.do())
}
