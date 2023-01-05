package ch4

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type IssueSearchResult struct {
	TotalCount int `json:total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

const IssuesURL = "https://api.github.com/search/issues"

// the command to run:
//go run ./... repo:golang/go is:open

func GithubIssues() {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		fmt.Printf("%d \t %s \t %s\n", item.Number, item.User.Login, item.Title)
	}
}

func searchIssues(searchQueries []string) (*IssueSearchResult, error) {
	var result IssueSearchResult
	query := url.QueryEscape(strings.Join(searchQueries, " "))
	response, err := http.Get(IssuesURL + "?q=" + query)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		err := response.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("error durig response.Body.Close: %w", err)
		}
		return nil, fmt.Errorf("search querry failed: %s", response.Status)
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		err := response.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("error durig response.Body.Close: %w", err)
		}
		return nil, fmt.Errorf("error during json Decode: %w", err)
	}
	return &result, nil
}
