package data_structure

import "fmt"

/*
*

LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
*/

/*
*
实现方案：
1. 数据结构 map + linkedlist
2. 实现算法
，每次增加
 1. 先看是否存在，如果存在，就直接移动就行了
 2. 再看容量，如果超过了，需要淘汰链表尾，需要注意的是，删除的key，应该是最后一个节点的rmKey，不是put的key
 3. 头插

，每次查询
 1. 看先是否存在，如果存在，移动到头
*/
type LRUCache struct {
	mp       map[int]*ListNode
	list     *LinkedList
	capacity int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{}
	cache.mp = make(map[int]*ListNode)
	cache.list = LinkedListNew()
	cache.capacity = capacity
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, present := this.mp[key]
	if !present {
		return -1
	}
	this.list.MoveToHead(node)
	fmt.Println("======================")
	this.list.Print(node)
	this.PrintMap()
	return node.Val
}
func (this *LRUCache) Put(key int, value int) {
	// 0. 先看是否存在，如果存在，直接更新值就行了
	if this.Get(key) != -1 {
		node, _ := this.mp[key]
		node.Val = value
		return
	}

	n := len(this.mp)
	// 1. 如果满了，首先删除对应的key
	if n+1 > this.capacity {
		rmNode := this.list.Tail.Prev
		rmkey := rmNode.Key
		// 删除key和链表中的节点
		delete(this.mp, rmkey)
		this.list.Remove(rmNode)
	}
	// 添加key和链表中的节点
	node := this.list.AddValue(key, value)
	this.mp[key] = node

	fmt.Println("======================")
	this.list.Print(node)
	this.PrintMap()
}

type LinkedList struct {
	Head, Tail *ListNode
}

// 工厂方法，构建一个链表
func LinkedListNew() *LinkedList {
	list := new(LinkedList)
	list.Head = new(ListNode)
	list.Tail = new(ListNode)
	list.Head.Next = list.Tail
	list.Tail.Prev = list.Head
	return list
}

func (this *LinkedList) AddValue(key, value int) *ListNode {
	node := new(ListNode)
	node.Val = value
	node.Key = key
	this.Add(node)
	return node
}

func (this *LinkedList) Add(node *ListNode) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next = node
	node.Next.Prev = node
}

func (this *LinkedList) Remove(node *ListNode) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

func (this *LinkedList) MoveToHead(node *ListNode) {
	this.Remove(node)
	this.Add(node)
}

func (this *LinkedList) Print(node *ListNode) {
	for p := this.Head.Next; p != this.Tail; p = p.Next {
		fmt.Print(p.Val, " ")
	}
	fmt.Println()
}

func (this *LRUCache) PrintMap() {
	for key, node := range this.mp {
		fmt.Print("[", key, ":", node.Val, "],")
	}
	fmt.Println()
}
