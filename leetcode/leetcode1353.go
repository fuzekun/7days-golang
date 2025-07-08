package leetcode

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
*
学习内容
1. 基础的流程语句，变量
2. 数组，切片
3.
*/
func maxEvents(events [][]int) int {
	n := len(events)
	maxDay := 0
	for _, events := range events {
		maxDay = max(events[1], maxDay)
	}
	fmt.Println("maxDay = ", maxDay)

	// 1. 排序
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	fmt.Println("按照第一维度排序数组为:", events)
	// 2. 遍历时间
	ans := 0
	pq := &IntHeap{}
	heap.Init(pq)

	for i, j := 1, 0; i <= maxDay; i++ {
		// 3.1 加入开头时间在当前时间之前的入堆
		for j < n && events[j][0] <= i {
			heap.Push(pq, events[j][1])
			j++
		}

		// 3.2 排除结尾时间小于当前的时间的，也就是排除已经结束的
		for pq.Len() > 0 && (*pq)[0] < i {
			heap.Pop(pq)
		}

		// 3.3 选择结尾时间最早的一个,今天开会
		if pq.Len() > 0 {
			top := heap.Pop(pq)
			fmt.Printf("第 %d 天选择结束时间为 %d 的会议\n", i, top)
			ans++
		}
	}
	return ans
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
