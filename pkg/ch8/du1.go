package ch8

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Du1() {

	roots := os.Args[2:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizesChan := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(fileSizesChan, root)
		}
		close(fileSizesChan)
	}()

	var numFiles, numBytes int64
	for size := range fileSizesChan {
		numFiles++
		numBytes += size
	}
	fmt.Printf("%d files \t %.1f GB\n", numFiles, float64(numBytes)/1e9)
}

func walkDir(fileSizesChan chan<- int64, root string) {
	for _, entry := range directoryEntries(root) {
		if entry.IsDir() {
			subDir := filepath.Join(root, entry.Name())
			walkDir(fileSizesChan, subDir)
		} else {
			fileSizesChan <- entry.Size()
		}
	}
}

func directoryEntries(root string) []os.FileInfo {
	entries, err := ioutil.ReadDir(root)
	if err != nil {
		log.Println("error from directory entries")
		log.Fatal(err)
	}
	return entries
}
