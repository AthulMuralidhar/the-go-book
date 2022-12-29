package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Ex1_7() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("error during fetch: %v\n", err)
			os.Exit(1)
		}
		//body, err := ioutil.ReadAll(response.Body)
		body, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			fmt.Printf("error during reading: %v\n", err)
			os.Exit(1)
		}
		err = response.Body.Close()
		if err != nil {
			fmt.Printf("error during response.Body.Close: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("given body: %s\n", body)
	}
}
