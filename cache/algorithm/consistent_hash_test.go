package algorithm

import (
	"fmt"
	"github.com/spf13/cast"
	"hash/crc32"
	"log"
	"strconv"
	"testing"
)

// 测试自定义哈希函数
func Test_ConsistentHash(t *testing.T) {
	// 基本使用
	m := NewConsistentHash(2, func(bytes []byte) uint32 {
		i, err := strconv.Atoi(string(bytes))
		if err != nil {
			log.Println("casting error: ", err)
			return 0
		}
		return cast.ToUint32(i)
	})
	m.Add("3", "7", "19", "8", "11")
	v := m.Get("10")
	fmt.Println(v)

	// 输出转换过程
	num := cast.ToString(0) + "8"
	bytes := []byte(num)
	hashUint := m.hash(bytes)
	fmt.Printf("num = %v, bytes = %v, hashUint = %v\n", num, bytes, hashUint)
}

// 测试CRC32哈希
func Test_test1(t *testing.T) {
	m := NewConsistentHash(3, nil)

	m.Add("3", "7", "19", "8", "11")
	m.Get("10")

	num := cast.ToString(0) + "8"
	bytes := []byte(num)
	hashUint := crc32.ChecksumIEEE(bytes)

	fmt.Println(hashUint)
}

func Test_binarySearch(t *testing.T) {
	index := make([]int, 0)
	// 测试不重复元素命中
	index = append(index, binarySearch([]int{1, 3, 5, 6}, 5))
	// 测试重复元素命中，最左边一个
	index = append(index, binarySearch([]int{1, 3, 5, 6, 8, 8, 8, 10}, 8))
	// 测试未命中，返回右边第一个
	index = append(index, binarySearch([]int{1, 3, 5, 16, 21, 27, 228, 1110}, 4))
	// 测试边界情况
	index = append(index, binarySearch([]int{1, 3, 5, 16, 21, 27, 228, 1110}, 1110))
	index = append(index, binarySearch([]int{1, 3, 5, 16, 21, 27, 228, 1110}, 1111))
	index = append(index, binarySearch([]int{1, 3, 5, 16, 21, 27, 228, 1110}, 1))
	index = append(index, binarySearch([]int{1, 3, 5, 16, 21, 27, 228, 1110}, -1))

	fmt.Printf("index: %v\n", index)
}
