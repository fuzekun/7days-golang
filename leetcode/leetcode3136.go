package leetcode

import "fmt"

/**

学习了
1. map的key不存在的情况下，如果是基础类型，会赋值成初始值，如果是指针类型会空指针
2. map的初始化方式：使用 := 就行了
*/

// 至少包含一个元音，至少包含一个辅音，至少包含一个数字，至少包含3个字符，至少有一个大写。不能包含除了字母和数字之外的其他字符
func isValid(word string) bool {

	intMap := map[byte]int{0: 1, 1: 2}
	fmt.Println(intMap[3])

	if len(word) < 3 {
		// 非正常长度
		return false
	}
	flagY, flagF := false, false
	chars := []rune(word)
	mp := map[rune]bool{'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	for _, c := range chars {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
			// 非正常字符
			return false
		}
		if mp[c] || mp[c-32] {
			// 元音
			flagY = true
		} else if c <= '0' || c >= '9' {
			// 非数字，只能是字母了；并且不是元音，所以只能是辅音了
			flagF = true
		}
	}
	return flagY && flagF
}
