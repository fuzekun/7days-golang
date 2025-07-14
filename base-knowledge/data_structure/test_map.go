package data_structure

import (
	"fmt"
	"sort"
)

// 测试map数组的初始化，直接引用的情况下，实际上是对数组的拷贝，不应该采用这种方式
func TestMapArrayInit() {
	// Version A: 正确的引用方式。
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: 错误的引用方式，item是值传递，也就是说是对数组的拷贝，而不是对数组的引用
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2                 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
}

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

// 测试数组排序
func TestMapSort() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	// 1. 把key拿出来
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	// 2. 对key进行排序
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	// 3. 遍历排序之后的key
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}

func TestRevMapKV() {
	// 加上了len，可以减少扩容，一次性分配，提升性能
	// 减少内存碎片，一次性分配足够的精确内存，扩容的话会引起内存的内部碎片的，因为有的内存分配了不会使用
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}

}
