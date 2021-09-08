package cache

import (
	"fmt"
	"log"
	"sync"
)

type Group struct {
	name      string
	getter    Getter
	cacheData cache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

/******************** kinds of api *******************/
/*****************************************************/

func NewGroup(name string, getter Getter, maxBytes uint64) *Group {
	if g, ok := groups[name]; ok {
		return g
	}
	// ensure that only one routine can create group in the same time
	mu.Lock()
	defer mu.Unlock()

	group := &Group{
		name:   name,
		getter: getter,
		cacheData: cache{
			maxBytes: maxBytes,
		},
	}
	groups[name] = group
	return group
}

func GetGroup(name string) *Group {
	mu.RLock()
	defer mu.RUnlock()

	if group, ok := groups[name]; ok {
		return group
	}
	return nil
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("invalid key")
	}

	mu.RLock()
	defer mu.RUnlock()
	if data, ok := g.cacheData.get(key); ok {
		log.Println("[Cache] hit")
		return data, nil
	}
	// get from local storage
	return g.getFromLocal(key)
}

func (g *Group) Add(key string, data ByteView) {
	mu.Lock()
	defer mu.Unlock()

	g.cacheData.add(key, data)
}

func (g *Group) getFromLocal(key string) (ByteView, error) {
	if g.getter == nil {
		return ByteView{}, fmt.Errorf("not found in cache, and no callback func")
	}

	data, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}

	log.Println("[Local Storage] hit")
	bytes := ByteView{
		data: data,
	}
	g.cacheData.add(key, bytes)
	return bytes, nil
}
