package ch4

import (
	"fmt"
	"log"
	"os"
)

type IssueSearchResult {
	
}

func GithubIssues() {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		fmt.Printf("%d \t %s \t %s\n"item.Number, item.User.Login, item.Title)
	}
}

func searchIssues(strings []string) (interface{}, interface{}) {
	
}
