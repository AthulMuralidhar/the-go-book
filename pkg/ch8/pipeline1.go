package ch8

import "fmt"

func Pipeline1() {
	intChannel := make(chan int)
	printChannel := make(chan int)
	go counter(intChannel)
	go squarer(intChannel, printChannel)

	for {
		x, ok := <-printChannel
		if !ok {
			break
		}
		fmt.Println(x)
	}
}

func counter(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
func squarer(srcCh, dstCh chan int) {
	for {
		x, ok := <-srcCh
		if !ok {
			break
		}
		dstCh <- x * x
	}
	close(dstCh)
}
