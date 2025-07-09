package leetcode

import (
	"errors"
	"fmt"
	"math"
	"time"
)

/*
*
1. make的基本语法，可以用来创建切片，映射，通道等等
2. rune是一个完整的unicode字符
3. 分支的语句, swich如何写，if...else如何写等等
4. abs和max, minInt，最好自己实现，math.Abs() 这个是float类型的

思维层面
实际上，可以简化为 Math.minInt(lr + ud + k * 2， i+1)
只要维护lr, ud两个变量就行了
*/
func maxDistance(ss string, k int) (int, error) {
	// 计算第一个用的时间
	start1 := time.Now()
	ans1 := solve1(ss, k)
	end1 := time.Now()
	time1 := end1.Sub(start1)

	// 计算第二个用的时间
	start2 := time.Now()
	ans2 := solve(ss, k)
	end2 := time.Now()
	time2 := end2.Sub(start2)

	if ans1 != ans2 {
		return -1, errors.New("两次计算结果不同，看看哪个方法错了")
	}

	fmt.Printf("两次计算的时间分别为 time1 = %s, time2 = %s\n", time1, time2)

	return ans1, nil
}

func solve1(ss string, k int) int {
	s := make([]int, 4)
	mp := map[rune]int{'N': 0, 'S': 1, 'W': 2, 'E': 3}
	ans := 0

	for _, c := range ss {
		dir := mp[c]
		s[dir]++

		lr := int(math.Abs(float64(s[0] - s[1])))
		ud := int(math.Abs(float64(s[2] - s[3])))

		lrChange := minInt(k, minInt(s[0], s[1]))
		restK := k - lrChange
		udChange := minInt(restK, minInt(s[2], s[3]))

		curValue := lr + ud + lrChange*2 + udChange*2
		if ans < curValue {
			ans = curValue
		}

		fmt.Println(s)
		fmt.Printf("lr = %d, ud = %d, lrChange = %d,"+
			" restK = %d, udChange = %d, curValue = %d\n",
			lr, ud, lrChange, restK, udChange, curValue)
	}
	return ans
}

func solve(s string, k int) int {
	latitude := 0
	longitude := 0
	ans := 0
	for i, c := range s {
		if c == 'N' {
			latitude++
		} else if c == 'S' {
			latitude--
		} else if c == 'E' {
			longitude++
		} else if c == 'W' {
			longitude--
		}
		ans = maxInt(ans, minInt(abs(latitude)+abs(longitude)+k*2, i+1))
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
