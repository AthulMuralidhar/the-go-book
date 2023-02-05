package bank2

import (
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	// Deposit [1..1000] concurrently.
	var n sync.WaitGroup
	for i := 1; i <= 10; i++ {
		n.Add(1)
		go func(amount int) {
			Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := Balance(), (10+1)*10/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
