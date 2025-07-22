package leetcode

func maximumUniqueSubarray(nums []int) int {
	ans, sum := 0, 0
	st := make(map[int]bool)
	n := len(nums)
	for i, j := 0, 0; j < n; j++ {
		x := nums[j]
		sum += x
		for st[x] {
			sum -= nums[i]
			delete(st, nums[i])
			i++
		}
		st[x] = true
		ans = max(ans, sum)
	}
	return ans
}
