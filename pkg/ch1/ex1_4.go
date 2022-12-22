package ch1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Ex1_4() {
	lineCounts := make(map[string]int)
	duplicateFiles := make(map[string]bool)

	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			err := fmt.Errorf("error from Dup3 during ioutil,ReadFile: %w", err)
			fmt.Println(err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if val, ok := lineCounts[line]; ok {
				lineCounts[line]++
				if val > 1 {
					duplicateFiles[file] = true
				}
			} else {
				lineCounts[line] = 1
			}
		}
	}
	for f, duplicate := range duplicateFiles {
		fmt.Printf("filename: %s \t duplicate: %t\n", f, duplicate)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	for line, num := range lineCounts {
		fmt.Printf("count: %d\t line: %s\n", num, line)
	}

}
