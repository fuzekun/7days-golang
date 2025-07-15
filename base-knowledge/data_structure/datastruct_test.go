package data_structure

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 0)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}

func TestLRUCache2(t *testing.T) {
	cache := Constructor(2)
	fmt.Println("ans = ", cache.Get(2))
	cache.Put(2, 6)
	fmt.Println("ans = ", cache.Get(1))
	// 注意这个修改了缓存中的内容，如果没有先看缓存中是否存在的情况，可能会出现将缓存挤出去的情况（容量此时满了的情况下）
	cache.Put(1, 5)
	cache.Put(1, 2)
	fmt.Println("ans = ", cache.Get(1))
	fmt.Println("ans = ", cache.Get(2))

}
