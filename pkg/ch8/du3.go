package ch8

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func Du3() {
	roots := os.Args[2:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	semaControlChan := make(chan int, 20) // making it a buffered channel allows only 20 files to be open at once
	fileSizeChan := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir3(root, &wg, fileSizeChan, semaControlChan)
	}

	go func() {
		wg.Wait()
		close(fileSizeChan)
	}()
	var tick <-chan time.Time
	tick = time.Tick(500 * time.Millisecond)
	var numFiles, numBytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizeChan:
			if !ok {
				break loop
			}
			numFiles++
			numBytes += size
		case <-tick:
			log.Printf("%d files \t %.1f B\n", numFiles, float64(numBytes))
		}

	}

	fmt.Printf("\n final totals: \n %d files \t %.1f GB\n", numFiles, float64(numBytes)/1e9)

}

func walkDir3(root string, wg *sync.WaitGroup, sizeChan chan<- int64, controlChanel chan int) {
	defer wg.Done()
	for _, entry := range directoryEntries3(root, controlChanel) {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(root, entry.Name())
			go walkDir3(subDir, wg, sizeChan, controlChanel)
		} else {
			sizeChan <- entry.Size()
		}
	}
}

func directoryEntries3(root string, controlChan chan int) []os.FileInfo {
	// need to have counting semaphores cause or else it opens too many files
	controlChan <- 0
	defer func() { <-controlChan }()
	entries, err := ioutil.ReadDir(root)
	if err != nil {
		log.Println("error from directory entries")
		log.Fatal(err)
	}
	return entries
}
