package ch8

import (
	"log"
	"os"
)

func Crawl3() {
	workList := make(chan []string)
	unseenList := make(chan string)
	go func() {

		workList <- os.Args[2:]
		log.Println("args sent to worklist")
		log.Printf("args: %s\n", os.Args[2:])
	}()
	for i := 0; i < 20; i++ {
		go func() { // the crawler go routine
			for link := range unseenList {
				foundLinks := crawl(link)
				go func() {
					workList <- foundLinks
					log.Println("foundlinks sent to worklist")
					log.Printf("foundlinks: %s\n", foundLinks)
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenList <- link // this is passsed to the crawler go routine
				log.Println("link sent to unseenlist")
				log.Printf("link: %s\n", link)
			}
		}
	}
}
