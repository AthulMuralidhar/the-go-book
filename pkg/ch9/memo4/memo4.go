package memo4

import "sync"

type result struct {
	value interface{}
	err   error
}

type Entry struct {
	result result
	ready  chan int
}

type Func func(key string) (interface{}, error)

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*Entry
}

// main focus here is duplicate suppression: meaning that the ready chan is used to broadcast to
// all the other go routines that the value is set in the entry set

func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	entry := m.cache[key]
	if entry == nil {
		// meaning that the result is not set
		entry = &Entry{
			ready: make(chan int),
		}

		m.cache[key] = entry
		m.mu.Unlock()
		entry.result.value, entry.result.err = m.f(key)
		close(entry.ready)
	} else {
		// this is just a read from key - no setting here
		m.mu.Unlock()
		// wait for ready to be closed,
		// the read from ready below is blocked until the channel is closed in line 38
		<-entry.ready
	}
	return entry.result.value, entry.result.err
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*Entry)}
}
