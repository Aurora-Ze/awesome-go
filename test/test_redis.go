package test

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/go-redis/redismock/v7"
	jsoniter "github.com/json-iterator/go"
)

type Use struct {
	Name string
	Id   int
}

func UseRedis() {

	cacheClient, cacheMock := GetCacheMock()

	// 模拟写
	//cacheMock.ExpectSet("kkey", "haha", time.Minute).SetVal("haha")

	// 模拟批量读
	use := &Use{Name: "test", Id: 1}
	useStr, err1 := jsoniter.MarshalToString(use)
	cacheMock.ExpectMGet("user:1:at:1").SetVal([]interface{}{useStr})

	defer cacheMock.ClearExpect()
	//obj := make([]interface{}, 0)
	//obj = append(obj, useStr)

	value1, err2 := cacheClient.MGet("user:1:at:1").Result()
	fmt.Printf("value: %v, err: %v, %v\n", value1, err1, err2)

	value2, err3 := cacheClient.MGet("user:1:at:1").Result()
	fmt.Printf("value2: %v, err: %v\n", value2, err3)

}

func GetCacheMock() (*redis.Client, redismock.ClientMock) {
	return redismock.NewClientMock()
}
