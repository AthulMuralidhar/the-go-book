package ch1

import (
	"fmt"
	"os"
)

func Echo2() {
	var s string
	const separator = " "
	argsArray := os.Args[1:]

	if CheckIfTesting() {
		argsArray = []string{TestingEnv}
	}

	for _, arg := range argsArray {
		s += arg + separator
	}
	if !CheckIfTesting() {
		fmt.Println("=== Echo2 ===")
		fmt.Printf("returned string: %s", s)
		fmt.Println("=== === ===")
	}

}
