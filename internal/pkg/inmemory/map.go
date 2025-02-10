package inmemory

import (
	"strings"
	"sync"
)

type InMemory struct {
	mutex sync.Mutex
	data  map[string][]byte
}

func New() *InMemory {
	return &InMemory{
		data: make(map[string][]byte),
	}
}

func (c *InMemory) GetByPrefix(prefix string) (value [][]byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for key, val := range c.data {
		if strings.HasPrefix(key, prefix) {
			value = append(value, val)
		}
	}
	return value
}

func (c *InMemory) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value
}

func (c *InMemory) Get(key string) (value []byte, ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, ok = c.data[key]

	return value, ok
}
