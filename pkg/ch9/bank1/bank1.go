// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank1

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(i int) {
	deposits <- i
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amt := <-deposits:
			balance += amt
		case balances <- balance:

		}
	}
}

func init() {
	go teller()
}
