package ch4

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

type IssueSearchResult {
	TotalCount int `json: "total_count"`
	Items []*Issue
}

type Issue struct {
Number int
HTMLURL string `json: "html_url"`
Title string
State string
User *User
CreatedAt time.Time `json: "created_at"`
Body string
}

type User struct {
	Login string
	HTMLURL string `json: "html_url"`
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

func searchIssues(searchQueries []string) (*IssueSearchResult, error) {
	query := url.QueryEscape()
}
