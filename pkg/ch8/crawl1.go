package ch8

import (
	"fmt"
	"github.com/AthulMuralidhar/the-go-book/pkg/ch5"
	"log"
	"os"
)

func Crawl1() {
	workList := os.Args[2:] // so not defining this leads to a forever loop :/
	var result []string
	seen := make(map[string]bool)
	if len(workList) > 0 {
		for _, item := range workList {
			if !seen[item] {
				seen[item] = true
				result = append(result, crawl(item)...)
			}

		}
	}
	for _, s := range result {
		fmt.Printf("linK: \t %s \n", s)
	}
}

func crawl(item string) []string {
	log.Printf("url: %s\n", item)
	list, err := ch5.Extract(item)
	if err != nil {
		log.Fatal(err)
	}
	return list
}
