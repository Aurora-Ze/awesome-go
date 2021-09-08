package lru

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

type V string

func (v V) Len() uint64 {
	return cast.ToUint64(len(v))
}

func Test_Cache(t *testing.T) {
	out := make(map[string]Value)
	// test initialize
	cache := NewCache(uint64(20), func(key string, value Value) {
		out[key] = value
	})
	// test get
	cache.Get("123")
	// test add
	cache.Add("123", V("value1"))
	cache.Add("haha", V("value2"))
	cache.Add("what", V("value3"))
	cache.Add("why", V("value4"))

	fmt.Printf("cache cur size = %d\n", cache.curBytes)

	fmt.Println("----------------------")
	fmt.Println(out["123"])
	fmt.Println(out["haha"])
}
