package memo1

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func Test_Memo1_Sequential(t *testing.T) {
	m := New(HttpGetBody)
	for url := range IncomingURLS() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func TestMemo_Get_Concurrent(t *testing.T) {

	m := New(HttpGetBody)
	var n sync.WaitGroup
	for url := range IncomingURLS() {
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
