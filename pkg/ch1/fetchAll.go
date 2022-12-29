package ch1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func FetchAll() {
	start := time.Now()
	stringChannel := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, stringChannel)
	}

	for _, _ = range os.Args[1:] {
		fmt.Println(<-stringChannel)
	}
	fmt.Printf("time elapsed from fetchAll: %f\n", time.Since(start).Seconds())
}

func fetch(url string, channel chan<- string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		channel <- fmt.Sprintf("error during http.Get: %v", err)
		return
	}
	numberOfBytes, err := io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		channel <- fmt.Sprintf("error during io.Copy: %v", err)
		return
	}
	err = response.Body.Close()
	if err != nil {
		channel <- fmt.Sprintf("error during response.Body.Close: %v", err)
		return
	}
	channel <- fmt.Sprintf("printing from fetch: %.2f secs \t %7d \t url:  %s\n", time.Since(start).Seconds(), numberOfBytes, url)
}
