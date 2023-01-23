package ch8

import (
	"fmt"
	"log"
	"os"
)

type Item struct {
	Error    error
	FileName string
}

func Thumbnail5() {
	fileNames := os.Args[2:]
	itemChan := make(chan Item, len(fileNames))
	for _, filename := range fileNames {
		go handleImangeFile(itemChan, filename)
	}
	for range fileNames {
		it := <-itemChan
		fmt.Printf("name of file: %s \n", it.FileName)
		if err := it.Error; err != nil {
			log.Fatal(err)
		}
	}
}

func handleImangeFile(itemChan chan<- Item, filename string) {
	it := Item{}
	f, err := ImageFile(filename)
	it.FileName = f
	if err != nil {
		it.Error = err
	}
	itemChan <- it
}

// so if i make two channels here for both file name and err, i cannot run them as they block the remaining runs
// however, i can make a struct type and carry around stuff in a single buffered channel with no worries
// seems pretty interesting :)
