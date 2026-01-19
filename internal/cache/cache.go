package cache

import (
	"sync"

	"github.com/diyliv/p2p/internal/models/interfaces"
)

type Cache struct {
	items map[interface{}]interface{}
	mutex sync.RWMutex
}

func NewCache() interfaces.Cache {
	return &Cache{
		items: make(map[interface{}]interface{}),
	}
}

func (c *Cache) Add(key interface{}, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items[key] = value
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, ok := c.items[key]
	return value, ok
}

func (c *Cache) Remove(key interface{}) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, exists := c.items[key]
	if exists {
		delete(c.items, key)
	}
	return exists
}

func (c *Cache) GetAll() interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	copy := make(map[interface{}]interface{}, len(c.items))
	for k, v := range c.items {
		copy[k] = v
	}
	return copy
}

func (c *Cache) GetAllAddresses() []string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	addrs := make([]string, 0, len(c.items))
	for key := range c.items {
		if addr, ok := key.(string); ok {
			addrs = append(addrs, addr)
		}
	}
	return addrs
}

func (c *Cache) HasAny() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.items) > 0
}

func (c *Cache) Count() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.items)
}

func (c *Cache) ForEach(fn func(key interface{}, value interface{})) {
	c.mutex.RLock()
	defer c.mutex.RLock()
	for key, value := range c.items {
		fn(key, value)
	}
}
