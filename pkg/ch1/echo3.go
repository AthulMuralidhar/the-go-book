package ch1

import (
	"fmt"
	"os"
	"strings"
)

func Echo3() {
	// as from the text of the book, we know that for every iteration of the loops in previous verison of echo
	// the variable s is made anew - i.e new memory location is cerated for each  "change" of the string
	// this here is a more momory efficient version of the same for loop where reourse use is critical

	argsArray := os.Args[1:]
	if CheckIfTesting() {
		argsArray = []string{TestingEnv}
	}

	s := strings.Join(argsArray, " ")

	if !CheckIfTesting() {
		fmt.Println("=== PowerEcho  ===")
		fmt.Printf("returned string: %s", s)
		fmt.Println("=== === ===")
		fmt.Println("just printing straing from strings.join")
		fmt.Println()
	}
}
