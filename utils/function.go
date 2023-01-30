package utils

import (
	"github.com/AthulMuralidhar/the-go-book/pkg/ch5"
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
	case f.String() == "Crawler" || f.String() == strings.ToLower("Crawler"):
		ch5.Crawler()
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
	case f.String() == "Thumbnail6" || f.String() == strings.ToLower("Thumbnail6"):
		ch8.Thumbnail6()
	case f.String() == "Crawl1" || f.String() == strings.ToLower("Crawl1"):
		ch8.Crawl1()
	case f.String() == "Crawl2" || f.String() == strings.ToLower("Crawl2"):
		ch8.Crawl2()
	case f.String() == "Crawl2b" || f.String() == strings.ToLower("Crawl2b"):
		ch8.Crawl2b()
	case f.String() == "Crawl3" || f.String() == strings.ToLower("Crawl3"):
		ch8.Crawl3()
	case f.String() == "Countdown1" || f.String() == strings.ToLower("Countdown1"):
		ch8.Countdown1()
	case f.String() == "Countdown2" || f.String() == strings.ToLower("Countdown2"):
		ch8.Countdown2()
	case f.String() == "Countdown3" || f.String() == strings.ToLower("Countdown3"):
		ch8.Countdown3()
	case f.String() == "Du1" || f.String() == strings.ToLower("Du1"):
		ch8.Du1()
	case f.String() == "Du2" || f.String() == strings.ToLower("Du2"):
		ch8.Du2()
	case f.String() == "Du3" || f.String() == strings.ToLower("Du3"):
		ch8.Du3()
	case f.String() == "Du4" || f.String() == strings.ToLower("Du4"):
		ch8.Du4()
	default:
		log.Fatal("name does not match existing functions")
	}
}
