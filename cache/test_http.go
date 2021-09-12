package cache

import (
	"fmt"
	"log"
)

// 模拟库数据
var storage = map[string]string{
	"user:zhangsan": "zhangsan:18",
	"user:lisi":     "lisi:20",
	"user:remiwu":   "remiwu:21",
	"user:aurora":   "aurora:21",
}

// 自定义读库
func getter(key string) ([]byte, error) {
	val, ok := storage[key]
	if !ok {
		return nil, fmt.Errorf("storage not found, key = %v\n", key)
	}
	bData := []byte(val)
	log.Printf("found, str = %v, byte = %v\n", val, bData)
	return bData, nil
}

// TestUsage 测试缓存使用
func TestUsage() {
	// 创建并添加缓存
	group := NewGroup("user_cache", GetterFunc(getter), 1000)
	group.Add("user:aurora", ByteView{
		data: []byte("aurora:22"),
	})

	// 启动HTTP服务
	StartServer()

	// 模拟数据不存在 group.Get("user:xxx")
	// 模拟读库      group.Get("user:zhangsan")
	// 模拟读缓存    group.Get("user:aurora")
}
