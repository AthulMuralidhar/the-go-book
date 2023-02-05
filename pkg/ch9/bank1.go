package ch9

import "fmt"

var deposits = make(chan int)
var withdrawls = make(chan int)

var balances = make(chan int)

func init() {
	go teller()
}

func teller() {
	var balance int
	for {
		select {
		case amt := <-deposits:
			balance += amt
		case withdrawls := <-withdrawls:
			balance -= withdrawls
		case balances <- balance:
			// i think this here must be used to update the balance variable, but
			//  that also means that thebalnce variable will be updated  by this channel
			// and can have a different value every time thiis channel is accessed -- leading to
			// multiple weites? isisnt that something we must avoud?

			// seems better to juust print the balances here
			// but looking at the tests, this print line seems superfolous indeed
			fmt.Println(<-balances)
		}

	}
}
