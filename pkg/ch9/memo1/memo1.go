package memo1

import (
	"io"
	"net/http"
	"sync"
)

type result struct {
	value interface{}
	err   error
}

type Func func(key string) (interface{}, error)

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	res, ok := m.cache[key]
	if !ok {
		res.value, res.err = m.f(key)
		m.cache[key] = res
	}
	return res.value, res.err
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func HttpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func IncomingURLS() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}
