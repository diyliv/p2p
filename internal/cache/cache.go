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

func (c *Cache) Get(key interface{}) (interface{}, bool)

func (c *Cache) Remove(key interface{}) bool

func (c *Cache) GetAll() interface{}

func (c *Cache) GetAllAddresses() []string

func (c *Cache) HasAny() bool

func (c *Cache) Count() int

func (c *Cache) ForEach(fn func(key interface{}, value interface{}))
