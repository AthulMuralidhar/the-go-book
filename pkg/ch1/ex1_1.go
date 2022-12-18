package ch1

import (
	"fmt"
	"os"
	"strings"
)

func Ex1_1() {
	fmt.Println(strings.Join(os.Args, " "))
}
