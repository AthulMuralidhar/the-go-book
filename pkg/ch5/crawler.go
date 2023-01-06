package ch5

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Crawler() {
	err := breadthFirst(crawl, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func crawl(url string) []string {
	fmt.Printf("url: \t %s", url)
	list, err := extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func extract(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			
		}
	}(response.Body)

}

func breadthFirst(f func(item string) []string, worklist []string) error {
	var result []string
	seen := make(map[string]bool)
	if len(worklist) > 0 {
		for _, item := range worklist {
			if !seen[item] {
				seen[item] = true
				result = append(result, f(item)...)
			}

		}
	}
}
