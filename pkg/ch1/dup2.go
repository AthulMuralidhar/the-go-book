package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func Dup2() {
	lineCounts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		err := countLines(os.Stdin, lineCounts)
		if err != nil {
			fmt.Printf("error from countLines: %v", err)
			return
		}
	} else {
		err := handleFileCountLines(files, lineCounts)
		if err != nil {
			fmt.Printf("error from handleFileCountLines: %v", err)
			return
		}
	}

	for line, num := range lineCounts {
		fmt.Printf("count: %d\t line: %s\n", num, line)
	}
}

func countLines(file *os.File, counts map[string]int) error {
	input := bufio.NewScanner(file)
	for input.Scan() {
		if _, ok := counts[input.Text()]; ok {
			counts[input.Text()]++
		} else {
			counts[input.Text()] = 1
		}
	}
	err := input.Err()
	if err != nil {
		return err
	}
	return nil
}

func handleFileCountLines(files []string, counts map[string]int) error {
	for _, file := range files {
		fOpened, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		err = countLines(fOpened, counts)
		if err != nil {
			fmt.Println("error during countLines, exiting...")
			return err
		}
		fOpened.Close()
	}
	return nil
}
