package data_structure

import (
	"container/heap"
	"fmt"
)

// IntHeap 是一个实现了 heap.Interface 的整数切片
type IntHeap []int

// Len 返回堆的长度
func (h IntHeap) Len() int { return len(h) }

// Less 比较元素大小，决定优先级（这里是最小堆，小的元素优先）
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap 交换元素位置
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push 添加元素到堆
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 移除并返回堆顶元素
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	// 注意这里应该干掉最后一个元素，因为heap的Pop已经帮你交换了old[0]和old[n-1]了
	// heap.Pop()实现了：1. 交换第一个和最后一个元素。2. 调整堆。所以你只需要干掉最后一个元素，其他的操作由heap帮你实现了。
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 示例使用
func TestIntHeap() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h) // 初始化堆

	heap.Push(h, 3)                   // 插入元素
	fmt.Printf("堆顶元素: %d\n", (*h)[0]) // 查看堆顶元素

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h)) // 按优先级弹出元素
	}
	// 输出: 1 2 3 5
}
