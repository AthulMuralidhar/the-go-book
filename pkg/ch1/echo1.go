package ch1

import (
	"fmt"
	"os"
)

func Echo1() {
	var s string
	const separator = " "

	//for i, arg := range os.Args {
	//	if i == 0 {
	//		continue
	//	}
	//	s += arg + separator
	//}

	for i := 1; i < len(os.Args); i++ { // start from 1 because we wanna skip the name of the function
		s += os.Args[i] + separator
	}

	fmt.Println("=== Echo1 ===")
	fmt.Printf("returned string: %s", s)
	fmt.Println("=== === ===")
}
