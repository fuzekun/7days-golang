package data_structure

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	ans := 0
	for p := head; p != nil; p = p.Next {
		ans = ans<<1 + p.Val
	}
	return ans
}

func TestLinkedList() {
	nums := []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}
	head := new(ListNode)
	p := head
	n := len(nums)
	for i, num := range nums {
		p.Val = num
		if i != n-1 {
			p.Next = new(ListNode)
		}
		p = p.Next
	}
	for p := head; p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
	fmt.Println("")
	ans := getDecimalValue(head)
	fmt.Println(ans)
}
