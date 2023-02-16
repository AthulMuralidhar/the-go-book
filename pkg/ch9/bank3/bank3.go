package bank3

import "sync"

var mu sync.Mutex
var balance int

func Deposit(amt int) {
	mu.Lock()
	balance += amt
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
