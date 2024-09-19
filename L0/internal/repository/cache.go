package repository

import (
	"sync"
)

type Cache struct {
	//sync.Mutex
	sync.RWMutex
	data map[string]Order
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]Order),
	}
}

func (c *Cache) Get(key string) (Order, bool) {
	c.RLock()
	defer c.RUnlock()
	order, exists := c.data[key]
	return order, exists
}

func (c *Cache) Set(key string, order Order) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = order
}
