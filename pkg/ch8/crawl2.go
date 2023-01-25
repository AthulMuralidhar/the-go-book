package ch8

import (
	"fmt"
	"os"
)

func Crawl2() {
	workList := make(chan []string)
	go func() {
		workList <- os.Args[2:]
	}()
	var result []string
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
	for _, s := range result {
		fmt.Printf("linK: \t %s \n", s)
	}
}
