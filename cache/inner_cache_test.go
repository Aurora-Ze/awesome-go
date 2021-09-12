package cache

import (
	"awesome-go/cache/lru"
	"fmt"
	"testing"
)

func Test_Cache(t *testing.T) {
	out := make(map[string]lru.Value)
	// test initialize
	cache := lru.NewCache(uint64(20), func(key string, value lru.Value) {
		out[key] = value
	})
	// test add and get
	cache.Add("user:aurora", NewByteView([]byte("aurora:22")))
	value, find := cache.Get("user:aurora")
	kv, ok := value.(ByteView)

	fmt.Printf("value = %v, find = %v\n", value, find)
	fmt.Printf("kv = %v, ok = %v\n", kv, ok)
	fmt.Printf("cache cur size = %d\n", cache.GetCurrentBytes())
}
