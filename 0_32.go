//map & make() & delete()
package main
import "fmt"

func main() {
	m := make(map[string]int)
	m2 := map[string]int{
	"James": 42,
	"Amy": 24}
	
	m["James"] = 42
	m["Amy"] = 24

	fmt.Println(m["James"])
	delete(m, "James")
	m["Bob"] = 8
	fmt.Println(m)
	fmt.Println(m2)

	m3 := map[int]int{
	8: 42,
	2: 6,
	4: 9,
	5: 3} 
	delete(m3, 2)
	fmt.Println(m3[4],m3[5])
}


