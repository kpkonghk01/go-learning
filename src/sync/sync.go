package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

// WIP: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/sync#ive-seen-other-examples-where-the-sync.mutex-is-embedded-into-the-struct.
