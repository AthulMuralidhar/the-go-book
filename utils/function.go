package utils

import (
	"log"
	"strings"
)
import "github.com/AthulMuralidhar/the-go-book/pkg/ch8"

type Function struct {
	Name string
}

func (f Function) String() string {
	return f.Name
}
func (f Function) Run() {
	switch {
	case f.String() == "Reverb2" || f.String() == strings.ToLower("Reverb2"):
		ch8.Reverb2()
	case f.String() == "NetCat3" || f.String() == strings.ToLower("NetCat3"):
		ch8.NetCat3()
	case f.String() == "ex8_3" || f.String() == "ex8.3":
		ch8.Ex8_3()
	case f.String() == "Pipeline1" || f.String() == strings.ToLower("Pipeline1"):
		ch8.Pipeline1()
	case f.String() == "Pipeline2" || f.String() == strings.ToLower("Pipeline2"):
		ch8.Pipeline2()
	case f.String() == "Pipeline3" || f.String() == strings.ToLower("Pipeline3"):
		ch8.Pipeline3()
	case f.String() == "Thumbnail1" || f.String() == strings.ToLower("Thumbnail1"):
		ch8.Thumbnail1()
	case f.String() == "Thumbnail5" || f.String() == strings.ToLower("Thumbnail5"):
		ch8.Thumbnail5()

	default:
		log.Fatal("name does not match existing functions")
	}
}
