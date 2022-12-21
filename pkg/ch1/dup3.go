package ch1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Dup3() {
	lineCounts := make(map[string]int)

	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			err := fmt.Errorf("error from Dup3 during ioutil,ReadFile: %w", err)
			fmt.Println(err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if _, ok := lineCounts[line]; ok {
				lineCounts[line]++
			} else {
				lineCounts[line] = 1
			}
		}
	}
}
