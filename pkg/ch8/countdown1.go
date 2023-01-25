package ch8

import (
	"log"
	"time"
)

func Countdown1() {
	log.Println("commencing countdown")
	tick := time.Tick(1 * time.Second)
	for i := 10; i > 0; i-- {
		log.Println(i)
		<-tick
	}
}
