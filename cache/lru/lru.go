package lru

import (
	"awesome-go/cache/structure"
	"github.com/spf13/cast"
)

type HandleFunc func(key string, value Value)

// Cache use LRU to eliminating the cache
type Cache struct {
	ll    *structure.DeList
	cache map[string]*structure.Element

	maxBytes uint64
	curBytes uint64

	OnEvicted HandleFunc
}

// Value the value must count how much memory it has used
type Value interface {
	Len() uint64
}

type entry struct {
	key   string
	value Value
}

// NewCache returns a initialized cache structure
func NewCache(max uint64, onEvicted HandleFunc) *Cache {
	return &Cache{
		maxBytes:  max,
		OnEvicted: onEvicted,
		ll:        structure.NewList(),
		cache:     make(map[string]*structure.Element),
	}
}

func (c *Cache) Get(key string) (Value, bool) {
	element, ok1 := c.cache[key]
	if !ok1 {
		return nil, false
	}
	// type assert
	kv, ok2 := element.Value.(*entry)
	if !ok2 {
		return nil, false
	}
	// move to head
	c.ll.RemoveToLast(element)
	return kv.value, true
}

func (c *Cache) CacheOut() {
	rmEle := c.ll.GetFirst()
	if rmEle == nil {
		return
	}

	c.ll.Remove(rmEle)
	kv := rmEle.Value.(*entry)
	delete(c.cache, kv.key)
	// update memory usage
	c.curBytes -= cast.ToUint64(len(kv.key)) + kv.value.Len()
	// execute callback function
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

func (c *Cache) Add(key string, value Value) {
	ele, exist := c.cache[key]
	if exist {
		c.ll.RemoveToLast(ele)
		kv := ele.Value.(*entry)
		kv.value = value
		// update memory
		c.curBytes += value.Len() - kv.value.Len()
	} else {
		ety := &entry{key: key, value: value}
		ele = c.ll.AddLast(ety) // ele means the added element
		c.cache[key] = &structure.Element{
			Value: ele,
		}
		c.curBytes += cast.ToUint64(len(key)) + value.Len()
	}

	// judge out of memory
	for c.curBytes > c.maxBytes {
		c.CacheOut()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
