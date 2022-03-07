// Select with default (to avoid deadlocks)

package main

import (
	"fmt"
	"time"	
)

func evenSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i
		}
	}
	ch <- result
}

func squareSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		result += i * i
	}
	ch <- result
}

func main() {
	evenCh := make(chan int)
	sqCh := make(chan int)
	
	go evenSum(0, 100, evenCh)
	go squareSum(0, 100, sqCh)
	
	select {
		case x := <- evenCh:
			fmt.Println(x)
		case y := <- sqCh:
			fmt.Println(y)
		default:
			fmt.Println("ничего не доступно")
			time.Sleep(50*time.Millisecond)
	}
	
	/*
	select {
	case <- c1:
		fmt.Println("received from c1")
	case <- c2:
		fmt.Println("received from c2")	
	}
	*/

}

