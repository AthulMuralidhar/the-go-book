package bad_bank

import "fmt"

var balance int

func BadBank() {
	go func() {
		deposit(200)
		fmt.Printf("balance: %d\n", balance)
	}()
	go deposit(100)

	fmt.Printf("balance: %d\n", balance)
}

func deposit(i int) {
	balance = balance + i
}
