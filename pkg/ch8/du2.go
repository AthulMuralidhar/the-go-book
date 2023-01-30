package ch8

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Du2() {
	var tick <-chan time.Time
	tick = time.Tick(2 * time.Second)
	var numFiles, numBytes int64
	fileSizesChan := make(chan int64)
	roots := os.Args[2:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	go func() {
		for _, root := range roots {
			walkDir(fileSizesChan, root)
		}
		log.Println("channel closing")
		close(fileSizesChan)
	}()
loop:
	for {
		select {
		case size, ok := <-fileSizesChan:
			if !ok {
				break loop
			}
			numFiles++
			numBytes += size
		case <-tick:
			//fmt.Printf("%d files \t %.1f GB\n", numFiles, float64(numBytes)/1e9)
			log.Printf("%d files \t %.1f B\n", numFiles, float64(numBytes))
		}

	}

	fmt.Printf("\n final totals: \n %d files \t %.1f GB\n", numFiles, float64(numBytes)/1e9)
}
