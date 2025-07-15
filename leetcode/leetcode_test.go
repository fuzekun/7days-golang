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

func TestSort(t *testing.T) {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	sortedArr := quickSort(arr)

	arr2 := []int{3, 6, 8, 10, 1, 2, 1}
	quickSort2(arr2)
	fmt.Println(sortedArr) // 输出: [1 1 2 3 6 8 10]
	fmt.Println(arr2)
	for i, cur := range sortedArr {
		if i != 0 && cur <= sortedArr[i-1] {
			t.Error("quickSort排序失败,排序后为:", sortedArr)
			break
		}
	}
	for i, cur := range arr2 {
		if i != 0 && cur < arr2[i-1] {
			t.Error("quickSort2排序失败，排序后的数组为:", quickSort(arr2))
		}
	}
}

func TestLeetcode3443(t *testing.T) {
	s := "NWSE"
	ans1, err := maxDistance(s, 1)

	if ans1 != 3 || err != nil {
		if err != nil {
			t.Error(err)
		} else {
			t.Error("测试用例1错误, ans = ", ans1)
		}
	}

	fmt.Println("========================================================")

	s = "NSWWEW"
	ans1, err = maxDistance(s, 3)
	if ans1 != 6 {
		if err != nil {
			t.Error(err)
		} else {
			t.Error("测试用例2错误, ans = ", ans1)
		}
	}
}

func Test_test(t *testing.T) {
	fmt.Println("hello world")
}

func Test_leetcode_3169(t *testing.T) {

	meetings := [][]int{{5, 7}, {1, 3}, {9, 10}}
	ans1 := countDays(10, meetings)

	fmt.Println(ans1)
}

func Test_leetcode_1760(t *testing.T) {
	nums := []int{431, 922, 158, 60, 192, 14, 788, 146, 788, 775, 772, 792, 68, 143, 376, 375, 877, 516, 595, 82, 56, 704, 160, 403, 713, 504, 67, 332, 26}
	ans := minimumSize(nums, 80)
	fmt.Println("ans = ", ans)
}

func Test_leetcode_1900(t *testing.T) {
	ans := earliestAndLatest(11, 2, 4)

	fmt.Println(ans)
}

func Test_leetcode_2410(t *testing.T) {
	ans := matchPlayersAndTrainers([]int{4, 7, 9}, []int{8, 2, 5, 8})
	fmt.Println(ans)
}

func Test_leetcode_3136(t *testing.T) {
	ans := isValid("234Adas")
	fmt.Println(ans)
	ans2 := isValid("a3$e")
	fmt.Println(ans2)
}
