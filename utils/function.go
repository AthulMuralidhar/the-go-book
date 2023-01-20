package utils

import "log"
import "github.com/AthulMuralidhar/the-go-book/pkg/ch8"

type Function struct {
	Name string
}

func (f Function) String() string {
	return f.Name
}
func (f Function) Run() {
	switch {
	case f.String() == "Reverb2":
		ch8.Reverb2()
	case f.String() == "NetCat3":
		ch8.NetCat3()
	case f.String() == "ex8_3" || f.String() == "ex8.3":
		ch8.Ex8_3()
	case f.String() == "Pipeline1":
		ch8.Pipeline1()
	case f.String() == "Pipeline2":
		ch8.Pipeline2()
	case f.String() == "Pipeline3":
		ch8.Pipeline3()
	case f.String() == "Thumbnail1":
		ch8.Thumbnail1()
	default:
		log.Fatal("name does not match existing functions")
	}
}
