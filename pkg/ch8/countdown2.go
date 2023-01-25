package ch8

import (
	"log"
	"os"
	"time"
)

func Countdown2() {
	abort := make(chan int)
	go func() {
		read, err := os.Stdin.Read(make([]byte, 1))
		if err != nil {
			log.Fatal(err)
		} // reads 1 byte
		abort <- read
	}()

	log.Println("commencing countdown")
	tick := time.Tick(1 * time.Second)
	for i := 10; i > 0; i-- {
		log.Println(i)
		<-tick
	}

	select {
	case <-abort:
		log.Println("launch aborted")
		return

	}

}
