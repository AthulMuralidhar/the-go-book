package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			// i see - the construction to the worker call here makes
			// sure that the waitgroup thing does not even have to be passed through
			// to the wworker funciton explicitrly
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}

func worker(i int) {
	fmt.Printf("worker %d starting\n", i)
	random := rand.Intn(10)

	time.Sleep(time.Second * time.Duration(random)) // sleep for a random time
	fmt.Printf("worker %d exiting\n", i)
}
