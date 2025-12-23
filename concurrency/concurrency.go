package concurrency

import (
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) Add(delta int64) {
	c.value += delta
}

func (c *Counter) Value() int64 {
	return c.value
}

func (c *Counter) AddAtomic(delta int64) {
	atomic.AddInt64(&c.value, delta)
}
