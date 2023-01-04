package ch4

import (
	"fmt"
	"log"
	"os"
)

func GithubIssues() {
	reult, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("issues: ")
}
