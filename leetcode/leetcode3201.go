package leetcode

/**
给你一个整数数组 nums。

nums 的子序列 sub 的长度为 x ，如果其满足以下条件，则称其为 有效子序列：

(sub[0] + sub[1]) % 2 == (sub[1] + sub[2]) % 2 == ... == (sub[x - 2] + sub[x - 1]) % 2
返回 nums 的 最长的有效子序列 的长度。

一个 子序列 指的是从原数组中删除一些元素（也可以不删除任何元素），剩余元素保持原来顺序组成的新数组。



示例 1：

输入： nums = [1,2,3,4]

输出： 4

解释：

最长的有效子序列是 [1, 2, 3, 4]。

示例 2：

输入： nums = [1,2,1,1,2,1,2]

输出： 6

解释：

最长的有效子序列是 [1, 2, 1, 2, 1, 2]。
[1,0,1,1,0,1,0]

示例 3：

输入： nums = [1,3]

输出： 2

解释：

最长的有效子序列是 [1, 3]。



提示：

2 <= nums.length <= 2 * 105
1 <= nums[i] <= 107
*/

func maximumLength(nums []int) int {
	// 0. 预处理，所有数字对2取余
	// 1. 因为是对2取余，和只有0和1
	// 2. 如果和是0的情况下，只有相邻完全相等
	// 3. 如果和是1的情况下，只有相邻完全不相等的长度
	max1, max2, max3, max4 := 0, 0, 0, 0
	preNum, preNum2 := 0, 1
	for _, num := range nums {
		x := num % 2
		if x == 0 {
			max1++
		} else {
			max2++
		}
		if preNum != x {
			preNum = x
			max3++
		}
		if preNum2 != x {
			preNum2 = x
			max4++
		}
	}
	return max(max1, max2, max3, max4)
}
