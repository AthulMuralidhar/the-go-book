package ch8

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func Thumbnail6() {
	var total int64
	fileNames := os.Args[2:]
	sizeChan := make(chan int64)
	var wg sync.WaitGroup
	for _, file := range fileNames {
		wg.Add(1)
		go handleImangeWithWG(&wg, sizeChan, file)
	}
	go func() {
		wg.Wait()
		close(sizeChan)
	}()

	for s := range sizeChan {
		total += s
	}
	fmt.Printf("size: %d bytes\n", total)
}

func handleImangeWithWG(wg *sync.WaitGroup, sizeChan chan<- int64, file string) {
	defer wg.Done()
	t, err := ImageFile(file)
	if err != nil {
		log.Fatal(err)
	}
	info, err := os.Stat(t)
	sizeChan <- info.Size()

}
