// Dynamic array with append method, function with variable arguments num, etc
package main

import "fmt"

func f(symbols ...string) {
	sum := 0
	for i, _ := range symbols {
		switch symbols[i]{
		case "w": sum += 3
		case "d": sum += 1
		}
	}
	fmt.Println(sum)
}

func main() {
	var last string
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	
	fmt.Scan(&last)
	results2 := append(results,last)
	f(results2...)
}
