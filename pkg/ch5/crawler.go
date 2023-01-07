package ch5

import (
	"fmt"
	"golang.org/x/net/html"
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
	fmt.Printf("target url: \t %s \n\n\n", url)
	list, err := extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func extract(url string) ([]string, error) {
	var links []string
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error during http.Get: response :%s \t status code: %s", response, response.StatusCode)
	}
	document, err := html.Parse(response.Body)

	if err != nil {
		return nil, err
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attribute := range n.Attr {
				if attribute.Key != "href" {
					continue
				}
				link, err := response.Request.URL.Parse(attribute.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(document, visitNode, nil)

	return links, nil
}

func forEachNode(node *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(node)
	}
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
	for _, s := range result {
		fmt.Printf("linK: \t %s \n", s)
	}
	return nil
}
