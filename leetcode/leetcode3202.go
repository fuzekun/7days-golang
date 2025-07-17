package leetcode

func leetcode3203MaximumLength(nums []int, k int) int {
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	res := 0
	for _, num := range nums {
		num %= k
		for prev := 0; prev < k; prev++ {
			// 以前一个开头，后一个结尾的 = 以后一个结尾，前一个开头的 + 1
			dp[prev][num] = dp[num][prev] + 1
			res = max(res, dp[prev][num])
		}
	}

	return res
}
