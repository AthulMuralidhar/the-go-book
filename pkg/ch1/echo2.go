package ch1

import (
	"fmt"
	"os"
)

func Echo2() {
	var s string
	const separator = " "
	for _, arg := range os.Args[1:] {
		s += arg + separator
	}

	fmt.Println("=== Echo2 ===")
	fmt.Printf("returned string: %s", s)
	fmt.Println("=== === ===")

}
