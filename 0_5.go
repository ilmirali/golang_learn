package main
import "fmt"

func main(){
	var i int 
	fmt.Scan(&i)
	
	if i%2 ==0 {
	fmt.Println("Even number")
	} else if i%2 != 0{
		fmt.Println("Odd number")
		}
}
