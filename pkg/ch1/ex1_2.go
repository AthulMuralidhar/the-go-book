package ch1

import (
	"fmt"
	"os"
)

func Ex1_2() {

	for i, arg := range os.Args {
		//fmt.Println("index %d has value %s", i, arg)
		// cannto do printf shit in print ln
		fmt.Printf("index %d has value %s", i, arg)
	}
}
