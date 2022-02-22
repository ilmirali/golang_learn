// method & pointer
package main
import "fmt"

type Account struct {
	balance int
}

func (acc *Account) deposit(x int) {
	acc.balance += x
}

func main() {
	a := Account{300}
	fmt.Println(a)
	a.deposit(100)
	fmt.Println(a)
}


