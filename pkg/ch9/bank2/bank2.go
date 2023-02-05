package bank2

import (
	"fmt"
	"github.com/AthulMuralidhar/the-go-book/pkg/ch9/bad_bank"
	"log"
)

var sema = make(chan struct{}, 1)
var balance2 int

func Deposit(i int) {
	log.Println("executing deposit go routine")
	sema <- struct{}{} // sending to tthe sema channel blocks the channel as it is only has len == 1
	// no one can access the channel other than to read so effectively only this go routine is using it
	balance2 += i
	<-sema
	// now the token is reaad from tje sema channel and is then released making it usuable again
}

func Balance() int {
	sema <- struct{}{} // i cannot read to sema here without  line 10 occuring, because this go routine would not exectue
	// untill the token is released from the previous run -- effecively stopping multiple writes

	// a good practice to do is to assume the deposite go routine is running
	// then the only way it can unblock the channel is only after the release of the tokem in line 10
	// if that has not happened then line 15 would wait untill line 10 is happened
	// effectively there is
	b := balance2
	<-sema
	return b
}

func Bank2() {

	var amt int
	go func() {
		log.Println("printing from first  read go routine")
		amt = Balance()
	}()
	fmt.Println(amt)

	go bad_bank.deposit(10)

	go func() {
		log.Println("printing from second  read go routine")
		amt = Balance()
	}()
	fmt.Println(amt)
}
