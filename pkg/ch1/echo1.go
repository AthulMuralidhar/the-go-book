package ch1

import (
	"fmt"
	"os"
)

func Echo1() {
	var s string
	const separator = " "

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		s += arg + separator
	}
	fmt.Println("=== Echo1 ===")
	fmt.Printf("returned string: %s", s)
	fmt.Println("=== === ===")
}
