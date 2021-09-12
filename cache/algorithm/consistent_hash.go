package algorithm

import (
	"github.com/spf13/cast"
	"hash/crc32"
	"sort"
)

type Hash func([]byte) uint32

// Map 一致性哈希的主要结构，用法可参照test文件中的示例
// 存储时每个真实节点都会根据 copyTimes 转化为若干个虚拟节点
type Map struct {
	hash      Hash
	copyTimes int
	key       []int          // 存放虚拟节点的哈希值，维持有序列表
	hashMap   map[int]string // 存放虚拟节点与真实节点的映射
}

// NewConsistentHash 创建一致性哈希的结构
// copyTimes 表示一个真实节点对应的虚拟节点个数
// hash 允许自定义哈希函数，传空则使用默认的CRC32
func NewConsistentHash(copyTimes int, hash Hash) *Map {
	m := &Map{
		hash:      hash,
		copyTimes: copyTimes,
		hashMap:   make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add 添加若干个真实节点，每个真实节点都会转化成多个虚拟节点进行存储
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.copyTimes; i++ {
			hashInt := cast.ToInt(m.hash([]byte(cast.ToString(i) + key)))
			m.key = append(m.key, hashInt)
			m.hashMap[hashInt] = key
		}
	}
	sort.Ints(m.key)
}

// Get 找到大于等于key哈希值的第一个虚拟节点，并返回它所对应的真实节点
func (m *Map) Get(key string) string {
	if len(key) == 0 {
		return ""
	}

	hashInt := cast.ToInt(m.hash([]byte(key)))
	idx := binarySearch(m.key, hashInt)

	v, _ := m.hashMap[m.key[idx%len(m.key)]]
	return v
}

// binarySearch 二分查找插入元素的位置，命中元素，返回最左边的下标（如果有重复的话）
// 边界情况：如果target最大或最小，则返回数组最后一个元素的下标，为了保证顺时针平均分配
func binarySearch(nums []int, target int) int {
	i, j := 0, len(nums)-1
	// 不在数组范围内则直接返回
	if target < nums[0] || target > nums[j] {
		return j
	}

	var mid int
	for i < j {
		mid = (i + j) >> 1
		if target > nums[mid] {
			i = mid + 1
		} else if target < nums[mid] {
			j = mid
		} else {
			j--
		}
	}
	return i
}
