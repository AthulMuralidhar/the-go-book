package bank2

// this program implements a binary semaphore ( a counting semaphore or a buffered channel with set length, that has len
// == 1
var balance int
var binarySemaphore = make(chan int, 1)

func Deposit(amt int) {
	binarySemaphore <- 0 // so this blocks the channel and no other program can access the next line until after line 11 is hit
	balance += amt
	<-binarySemaphore
}

func Balance() int {
	binarySemaphore <- 0
	b := balance
	<-binarySemaphore
	return b
}
