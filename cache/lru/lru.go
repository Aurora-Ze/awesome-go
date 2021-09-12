package lru

import (
	"awesome-go/cache/structure"
	"github.com/spf13/cast"
	"log"
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

func (e entry) Len() uint64 {
	return cast.ToUint64(len(e.key)) + e.value.Len()
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
	kv, ok2 := element.Pair.(*entry)
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
	kv, ok := rmEle.Pair.(*entry)
	if !ok {
		log.Println("casting element value error")
		return
	}

	delete(c.cache, kv.key)
	// update memory usage
	c.curBytes -= kv.Len()
	// execute callback function
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

// Add 添加缓存，以键值对的形式存储在 element 结构体的Pair字段中
func (c *Cache) Add(key string, value Value) {
	ele, exist := c.cache[key]
	if exist {
		c.ll.RemoveToLast(ele)
		// fixme
		kv, ok := ele.Pair.(*entry)
		if !ok {
			log.Println("casting element value error")
			return
		}
		old := kv.value
		kv.value = value
		ele.Pair = kv
		c.cache[key] = ele
		// update memory
		c.curBytes += value.Len() - old.Len()
	} else {
		addEntry := &entry{
			key:   key,
			value: value,
		}
		ele = c.ll.AddLast(addEntry) // ele means the added element
		c.cache[key] = ele
		c.curBytes += addEntry.Len()
	}

	// judge out of memory
	for c.curBytes > c.maxBytes {
		c.CacheOut()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}

func (c *Cache) GetMaxBytes() uint64 {
	return c.maxBytes
}

func (c *Cache) GetCurrentBytes() uint64 {
	return c.curBytes
}
