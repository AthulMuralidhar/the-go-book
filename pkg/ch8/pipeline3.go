package ch8

import "fmt"

func Pipeline3() {
	intChannel := make(chan int)
	printChannel := make(chan int)
	go counter3(intChannel)
	go squarer3(intChannel, printChannel)
	printer(printChannel) // so adding go to this line never prints :/
}

func printer(channel <-chan int) {
	for i := range channel {
		fmt.Println(i)
	}
}

func squarer3(src <-chan int, dst chan<- int) {
	for i := range src {
		dst <- i * i
	}
	close(dst)
}

func counter3(dstChannel chan<- int) {
	for i := 0; i < 100; i++ {
		dstChannel <- i
	}
	close(dstChannel)
}
