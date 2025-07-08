package leetcode

import (
	"fmt"
	"testing"
)

func Test_1353(t *testing.T) {
	events1 := [][]int{
		{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3},
	}

	events2 := [][]int{
		{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5},
	}
	fmt.Println("=================================")
	ans1 := maxEvents(events1)
	fmt.Println("ans1 = ", ans1)
	fmt.Println("=================================")
	ans2 := maxEvents(events2)
	fmt.Println("ans2 = ", ans2)
	if ans1 != 5 {
		t.Error("第一个测试用例失败,ans = ", ans1)
	}

	if ans2 != 5 {
		t.Error("第二个测试用例失败, ans2 = ", ans2)
	}
}
