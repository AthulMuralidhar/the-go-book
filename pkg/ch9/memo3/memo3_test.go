package memo3

import (
	"fmt"
	"github.com/AthulMuralidhar/the-go-book/pkg/ch9/memo1"
	"log"
	"sync"
	"testing"
	"time"
)

func TestMemo_Get_Concurrent(t *testing.T) {
	m := New(memo1.HttpGetBody)
	var n sync.WaitGroup
	for url := range memo1.IncomingURLS() {
		n.Add(1)

		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)

	}
	n.Wait()
}
