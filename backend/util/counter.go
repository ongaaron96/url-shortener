package util

import "sync"

type Counter struct {
	mutex sync.Mutex
	count uint64
}

func NewCounter(startCount uint64) *Counter {
	return &Counter{
		count: startCount,
	}
}

func (c *Counter) GetNextCount() uint64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count += 1
	return c.count
}
