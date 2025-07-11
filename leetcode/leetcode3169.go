package leetcode

import (
	"fmt"
	"sort"
)

/**
一、题目
给你一个正整数 days，表示员工可工作的总天数（从第 1 天开始）。
另给你一个二维数组 meetings，长度为 n，其中 meetings[i] = [start_i, end_i] 表示第 i 次会议的开始和结束天数（包含首尾）。
返回员工可工作且没有安排会议的天数。
注意：会议时间可能会有重叠。

示例 1：
输入：days = 10, meetings = [[5,7],[1,3],[9,10]]
输出：2
解释：
第 4 天和第 8 天没有安排会议。

示例 2：
输入：days = 5, meetings = [[2,4],[1,3]]
输出：1
解释：
第 5 天没有安排会议。

示例 3：
输入：days = 6, meetings = [[1,6]]
输出：0
解释：
所有工作日都安排了会议。
*/
/**
二、学习到的东西总结
语法
1. 数组（切片）排序
2. 循环，分支，变量，赋值等等
3. 测试用例应该怎么编写

逻辑
4. 数组中数字的个数。+1，-1还是不加不减呢？

*/
/**
思路：
1. 给meetings排个序，按照第一维度
2. 区间合并，如果出现了不连续的情况，就进行统计

*/
func countDays(days int, meetings [][]int) int {

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	ans := 0
	preEnd := meetings[0][1]
	for _, meeting := range meetings {
		fmt.Println(meeting)
		if preEnd < meeting[0] {
			fmt.Println("add : ", meeting[0]-preEnd)
			// 1. 如果出现了不连续的区间，答案应该加上两个连续区间的差值
			ans += meeting[0] - preEnd - 1
			preEnd = meeting[1]
		} else if preEnd < meeting[1] {
			// 2. 如果出现了连续的区间，进行区间合并
			preEnd = meeting[1]
		}
	}
	// 3. 加上开头和结尾
	fmt.Println("end : ", days-preEnd, ", begin:", meetings[0][0]-1)
	ans += days - preEnd + meetings[0][0] - 1
	return ans
}
