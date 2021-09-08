package cache

import (
	"awesome-go/cache/lru"
	"sync"
)

type cache struct {
	lru      *lru.Cache
	mu       sync.Mutex
	maxBytes uint64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// lazy initialize
	if c.lru == nil {
		c.lru = lru.NewCache(c.maxBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (ByteView, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru != nil {
		if kv, exist := c.lru.Get(key); exist {
			return kv.(ByteView), true
		}
	}
	return ByteView{}, false
}
