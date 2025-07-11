package leetcode

import (
	"fmt"
	"math"
)

/**
题目描述：
给你一个整数数组 nums ，其中 nums[i] 表示第i个袋子里球的数目。
同时给你一个整数 maxOperations 。你可以进行如下操作至多 maxOperations 次：
选择任意一个袋子，并将袋子里的球分到 2 个新的袋子中，每个袋子里都有正整数个球。
比方说，一个袋子里有5个球，你可以把它们分到两个新袋子里，分别有1个和4个球，或者分别有2个和3个球。
你的开销是单个袋子里球数目的最大值，你想要最小化开销。
请你返回进行上述操作后的最小开销。

最小化最大值，经典的二分算法



示例 1：
输入：nums = [9], maxOperations =2
输出：3
解释：
- 将装有9个球的袋子分成装有6个和 3 个球的袋子。[9] -> [6,3] 。
- 将装有6个球的袋子分成装有3个和3个球的袋子。[6,3] -> [3,3,3] 。
装有最多球的袋子里装有 3 个球，所以开销为 3 并返回3。

示例 2：

输入：nums = [2,4,8,2], maxOperations = 4
输出：2
解释：
- 将装有 8 个球的袋子分成装有 4 个和 4 个球的袋子。[2,4,8,2] -> [2,4,4,4,2] 。
- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,4,4,4,2] -> [2,2,2,4,4,2] 。
- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,2,2,4,4,2] -> [2,2,2,2,2,4,2] 。
- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,2,2,2,2,4,2] -> [2,2,2,2,2,2,2,2] 。
装有最多球的袋子里装有 2 个球，所以开销为 2 并返回 2 。

示例 3：
输入：nums = [7,17], maxOperations = 2
输出：7
3 - 1 = 2

*/
/**
学习的点：
1. int的最大值math.MaxInt32
*/

func minimumSize(nums []int, maxOperations int) int {
	l, r := 1, math.MaxInt32-1

	for l < r {
		mid := (l + r) / 2
		isOk := check(nums, mid, maxOperations)
		fmt.Printf("l = %d, r = %d, curValue = %d, isOk : %t\n", l, r, mid, isOk)
		if isOk {
			// 可以的情况下，可以缩小范围了
			r = mid
		} else {
			// 不行的情况下，需要把值扩大一下
			l = mid + 1
		}
	}
	return l
}

func check(nums []int, x int, maxOperations int) bool {
	cnt := 0
	for _, num := range nums {
		// 上取整-1
		cnt += num/x - 1
		if num%x != 0 {
			cnt++
		}
		if cnt > maxOperations {
			// int类型的可能会出现报精度的
			return false
		}
	}
	// 如果次数小于等于的情况下，说明是可行的
	fmt.Println("cnt = ", cnt)
	return cnt <= maxOperations
}
