package leetcode

import (
	"log"
	"math"
)

/**

n 名运动员参与一场锦标赛，所有运动员站成一排，并根据 最开始的 站位从 1 到 n 编号（运动员 1 是这一排中的第一个运动员，运动员 2 是第二个运动员，依此类推）。

锦标赛由多个回合组成（从回合 1 开始）。每一回合中，这一排从前往后数的第 i 名运动员需要与从后往前数的第 i 名运动员比拼，获胜者将会进入下一回合。如果当前回合中运动员数目为奇数，那么中间那位运动员将轮空晋级下一回合。

例如，当前回合中，运动员 1, 2, 4, 6, 7 站成一排
运动员 1 需要和运动员 7 比拼
运动员 2 需要和运动员 6 比拼
运动员 4 轮空晋级下一回合
每回合结束后，获胜者将会基于最开始分配给他们的原始顺序（升序）重新排成一排。

编号为 firstPlayer 和 secondPlayer 的运动员是本场锦标赛中的最佳运动员。在他们开始比拼之前，完全可以战胜任何其他运动员。而任意两个其他运动员进行比拼时，其中任意一个都有获胜的可能，因此你可以 裁定 谁是这一回合的获胜者。

给你三个整数 n、firstPlayer 和 secondPlayer 。返回一个由两个值组成的整数数组，分别表示两位最佳运动员在本场锦标赛中比拼的 最早 回合数和 最晚 回合数。



示例 1：

输入：n = 11, firstPlayer = 2, secondPlayer = 4
输出：[3,4]
解释：
一种能够产生最早回合数的情景是：
回合 1：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
回合 2：2, 3, 4, 5, 6, 11
回合 3：2, 3, 4
一种能够产生最晚回合数的情景是：
回合 1：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
回合 2：1, 2, 3, 4, 5, 6
回合 3：1, 2, 4
回合 4：2, 4
示例 2：

输入：n = 5, firstPlayer = 1, secondPlayer = 5
输出：[1,1]
解释：两名最佳运动员 1 和 5 将会在回合 1 进行比拼。
不存在使他们在其他回合进行比拼的可能。


提示：

2 <= n <= 28
1 <= firstPlayer < secondPlayer <= n
*/

/*
*
语法学习
1. go语言版的记忆化搜索
2. make创建数组，分配内存之后，数据的赋初值
3. 函数闭包
4. 内部函数

算法学习
1.
*/
var where = func() {
	log.SetFlags(log.Llongfile)
	log.Print()
}

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	where()
	type pair struct{ earliest, latest int }
	// 0.初始化记忆数组为最大值
	memo := make([][][]pair, n+1)
	for i := range memo {
		memo[i] = make([][]pair, n+1)
		for j := range memo[i] {
			memo[i][j] = make([]pair, n+1)
		}
	}

	var dfs func(int, int, int) pair
	dfs = func(n, first, second int) (res pair) {
		// 0.1 递归基
		if first+second == n+1 {
			return pair{1, 1}
		}

		// 0.2 保证A左边的人数比B右边的人数少
		if first+second > n+1 {
			first, second = n+1-second, n+1-first
		}

		// 1. 记忆化
		p := &memo[n][first][second]
		if p.earliest > 0 {
			return *p
		}
		defer func() { *p = res }()

		// 2. 递归求解
		// 2.1 计算剩下的人数
		m := (n + 1) / 2
		// 2.2 计算mid的范围
		minMid, maxMid := 0, second-first
		if second > m {
			// B在中轴线右侧
			minMid, maxMid = second-n/2-1, m-first
		}
		// 2.3 递归计算结果
		res.earliest = math.MaxInt
		for left := range first {
			for mid := minMid; mid < maxMid; mid++ {
				r := dfs(m, left+1, left+mid+2)
				res.earliest = min(res.earliest, r.earliest)
				res.latest = max(res.latest, r.latest)
			}
		}
		// 4 加上当前回合
		res.earliest++
		res.latest++
		return res
	}

	ans := dfs(n, firstPlayer, secondPlayer)
	return []int{ans.earliest, ans.latest}
}
