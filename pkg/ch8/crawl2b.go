package ch8

import (
	"fmt"
	"github.com/AthulMuralidhar/the-go-book/pkg/ch5"
	"log"
	"os"
)

// Crawl2b using counting semaphores here
func Crawl2b() {
	tokens := make(chan int, 20)
	var n int // n == work yet to be done
	workList := make(chan []string)
	n++
	go func() {
		workList <- os.Args[2:]
	}()
	var result []string
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					workList <- crawl2b(tokens, link)
				}(link)
			}
		}
	}

	for _, s := range result {
		fmt.Printf("linK: \t %s \n", s)
	}
}

func crawl2b(tokens chan int, link string) []string {
	log.Printf("url: %s\n", link)
	tokens <- 0 // add token
	list, err := ch5.Extract(link)
	<-tokens // release token
	if err != nil {
		log.Fatal(err)
	}
	return list
}
