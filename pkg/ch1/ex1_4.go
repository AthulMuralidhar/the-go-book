package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func Ex1_4() {
	lineCounts := make(map[string]int)
	files := os.Args[1:]
	duplicateFiles := make(map[string]bool)

	if len(files) == 0 {
		err := fmt.Errorf("error: we dont do without files now")
		fmt.Println(err)
	} else {
		for _, file := range files {
			fOpened, err := os.Open(file)
			if err != nil {
				fmt.Printf("dup2: %v\n", err)
				continue
			}
			input := bufio.NewScanner(os.Stdin)

			for input.Scan() {
				if _, ok := lineCounts[input.Text()]; ok {
					duplicateFiles[file] = true
					lineCounts[input.Text()]++
				} else {
					duplicateFiles[file] = false
					lineCounts[input.Text()] = 1
				}
				if input.Err() != nil {
					err := fmt.Errorf("error during input.Scan(): %w\n", input.Err())
					fmt.Println(err)
				}
			}
			err = fOpened.Close()
			if err != nil {
				err := fmt.Errorf("error during file.Close(): %w\n", input.Err())
				fmt.Println(err)
			}
		}
	}
	for f, duplicate := range duplicateFiles {
		fmt.Printf("filename: %s \t duplicate: %t", f, duplicate)
	}

}
