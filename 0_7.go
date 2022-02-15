package main
import "fmt"

func main(){
	var i int
	fmt.Println("Please, enter 0..9 or 9..19 or 20..30:")
	fmt.Scan(&i)
		
	switch {
	case i >0 && i<10: fmt.Println("0..10")
	case i>9  && i<20: fmt.Println("10..20")
	case i>19 && i<31: fmt.Println("20..30")
	default: fmt.Println("something else")
	}
}
