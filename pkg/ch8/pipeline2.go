package ch8

import "fmt"

func Pipeline2() {
	intChannel := make(chan int)
	printChannel := make(chan int)
	go counter2(intChannel)
	go squarer2(intChannel, printChannel)

	for i := range printChannel {
		fmt.Println(i)
	}
}

func squarer2(channel chan int, channel2 chan int) {
	for i := range channel {
		channel2 <- i * i
	}
	close(channel2)
}

func counter2(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i
	}
	close(channel) // no need to explicity close this, but file.close must be called regardless // these are 2 different things lol
}
