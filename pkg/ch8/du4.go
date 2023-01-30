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

func Du4() {
	roots := os.Args[2:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	semaControlChan := make(chan int, 20) // making it a buffered channel allows only 20 files to be open at once
	fileSizeChan := make(chan int64)
	// done chan is used to boradcast a close to all the channels and handle the draining gracefully
	doneChan := make(chan int)

	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir4(root, &wg, fileSizeChan, semaControlChan, doneChan)
	}

	go func() {
		wg.Wait()
		close(fileSizeChan)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		log.Println("aborting")
		close(doneChan)
	}()

	var tick <-chan time.Time
	tick = time.Tick(500 * time.Millisecond)
	var numFiles, numBytes int64
loop:
	for {
		select {
		case <-doneChan:
			for _ = range fileSizeChan {
				//log.Printf("draining filesizechan: %d", i)
			}
			return

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

func walkDir4(root string, wg *sync.WaitGroup, sizeChan chan int64, controlChan, doneChan chan int) {
	defer wg.Done()
	if cancelled(doneChan) {
		return
	}
	for _, entry := range directoryEntries4(root, controlChan, doneChan) {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(root, entry.Name())
			go walkDir4(subDir, wg, sizeChan, controlChan, doneChan)
		} else {
			sizeChan <- entry.Size()
		}
	}
}

func directoryEntries4(root string, controlChan, doneChan chan int) []os.FileInfo {
	select {
	case controlChan <- 0:
	case <-doneChan:
		return nil
	}
	defer func() { <-controlChan }()
	entries, err := ioutil.ReadDir(root)
	if err != nil {
		log.Println("error from directory entries")
		log.Fatal(err)
	}
	return entries
}

func cancelled(doneChan <-chan int) bool {
	select {
	case <-doneChan:
		return true
	default:
		return false
	}
}
