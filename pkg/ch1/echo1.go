package ch1

import (
	"fmt"
	"os"
)

const TestingEnv = "testing"

func Echo1() {
	var s string
	const separator = " "
	argsArray := os.Args

	//for i, arg := range os.Args {
	//	if i == 0 {
	//		continue
	//	}
	//	s += arg + separator
	//}

	if CheckIfTesting() {
		argsArray = []string{TestingEnv}
	}

	for i := 1; i < len(argsArray); i++ { // start from 1 because we wanna skip the name of the function
		s += os.Args[i] + separator
	}

	if !CheckIfTesting() {
		fmt.Println("=== Echo1 ===")
		fmt.Printf("returned string: %s", s)
		fmt.Println("=== === ===")
	}
}

func CheckIfTesting() bool {
	if os.Getenv("GO_ENV") == TestingEnv {
		return true
	}
	return false
}
