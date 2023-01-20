package main

import (
	"github.com/AthulMuralidhar/the-go-book/utils"
	"os"
)

func main() {
	arg1 := os.Args[1]
	f := utils.Function{Name: arg1}
	f.Run()
}
