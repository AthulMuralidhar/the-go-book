package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func Dup1() {
	lineCounts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if _, ok := lineCounts[input.Text()]; ok {
			lineCounts[input.Text()]++
		} else {
			lineCounts[input.Text()] = 1
		}
		if input.Err() != nil {
			_ = fmt.Errorf("error: %w\n", input.Err())
		}
	}

	for line, num := range lineCounts {
		fmt.Printf("%d \t %s \n", num, line)
	}
}
