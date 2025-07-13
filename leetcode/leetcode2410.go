package leetcode

import (
	"fmt"
	"sort"
)

/*
*

明显的一个二分图匹配，具体的算法就是尝试算法吧。dfs + 回溯
贪心好像也可以，最小的开始匹配，单指针从左到右找最小的
*/
func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	fmt.Println("players = ", players)
	fmt.Println("trainers = ", trainers)
	j := 0
	m := len(trainers)
	ans := 0
	for _, player := range players {
		for j < m && player > trainers[j] {
			j++
		}
		if j >= m {
			break
		}
		j++
		ans++
	}
	return ans
}
