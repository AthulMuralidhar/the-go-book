package main

import (
	"fmt"
	"github.com/AthulMuralidhar/the-go-book/pkg/ch9/memo1"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	m := memo1.New(httpGetBody)

	incomingURLS := os.Args[1:]
	for _, url := range incomingURLS {
		start := time.Now()
		val, err := m.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("url: %s \t time: %s \t bytes: %d \n", url, time.Since(start), len(val.([]byte)))
	}

}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return io.ReadAll(resp.Body)
}
