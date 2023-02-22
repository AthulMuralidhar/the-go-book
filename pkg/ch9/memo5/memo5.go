package memo5

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type Entry struct {
	result result
	ready  chan int
}

func (e *Entry) call(f Func, key string) {
	e.result.value, e.result.err = f(key)
	close(e.ready) // broadcast step
}

func (e *Entry) deliver(response chan<- result) {
	// same as memo4 - wait for entry to be ready
	<-e.ready
	// send to memo.req chan
	response <- e.result
}

type request struct {
	key      string
	response chan<- result
}

type Memo struct {
	requests chan request
}

func (m *Memo) server(f Func) {
	cache := make(map[string]*Entry)
	for req := range m.requests {
		entry := cache[req.key]
		if entry == nil {
			entry = &Entry{ready: make(chan int)}
			cache[req.key] = entry
			go entry.call(f, req.key)
		}
		go entry.deliver(req.response)
	}
}

func New(f Func) *Memo {
	m := &Memo{requests: make(chan request)}
	go m.server(f)
	return m
}
