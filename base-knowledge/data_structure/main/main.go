package main

import (
	"base-knowledge/data_structure"
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	// 测试数组定义、分配内存
	testArrayDefine()
	// 测试二维数组
	test2DemArray()
	// 测试切片引用
	testArraySlice()
	// 测试buffer数组
	testBuffer()
	// 测试切片的复制和追加
	testCopyAppend()
	//
	testAppendByte()
	// 测试字符串反转
	testReverse()
	// 测试链表
	data_structure.TestLinkedList()
	// 测试字符串的相关
	testStrings()
}

func testArrayDefine() {
	// 定义 + 分配内存
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	var arrX = [200]int{}
	// 定义变量
	var arry []int
	// 分配内存
	arry = make([]int, 100)
	arrX[0] = 1
	fmt.Println(arrLazy)
	fmt.Println(arrKeyValue)
	fmt.Println(arrX[0])
	fmt.Println(arry[99])
}

// 输出长度为3的数组
func fp(a *[3]int) { fmt.Println(a) }

func test2DemArray() {
	for i := 0; i < 3; i++ {
		// 调用的时候，传入数组的地址就行了，[3]int{x,y,z}创建了一个数组
		fp(&[3]int{i, i * i, i * i * i})
	}
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
		a[i] = 10
	}
	return s
}

func testArraySlice() {
	var arr = [5]int{0, 1, 2, 3, 4}
	fmt.Println("初始数组为:", arr)
	ans := sum(arr[:])
	fmt.Println(ans)
	fmt.Println("使用切片，调用后的数组为", arr)
	testCallArray(arr)
	fmt.Println("使用数值应用，调用后的数组为", arr)
}

func testCallArray(a [5]int) {
	for i, _ := range a {
		a[i] = 5
	}
}

func testBuffer() {
	buffer := bytes.Buffer{}
	buffer.WriteString("hello world")
	fmt.Println(buffer.String())
}

func testCopyAppend() {

	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(slTo, 4, 5, 6)
	fmt.Println(sl3)
}

func testAppendByte() {
	old := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	ans := data_structure.AppendByte(old, 1, 23, 4)
	fmt.Println(ans)
}

func reverse(s string) string {
	runes := []rune(s)
	n, h := len(runes), len(runes)/2
	for i := 0; i < h; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
func testReverse() {
	// reverse a string:
	str := "Google"
	sl := []byte(str)
	var rev [100]byte
	j := 0
	for i := len(sl) - 1; i >= 0; i-- {
		rev[j] = sl[i]
		j++
	}
	str_rev := string(rev[:])
	fmt.Printf("The reversed string is -%s-\n", str_rev)
	// variant: "in place" using swapping
	str2 := "Google"
	sl2 := []byte(str2)
	for i, j := 0, len(sl2)-1; i < j; i, j = i+1, j-1 {
		sl2[i], sl2[j] = sl2[j], sl2[i]
	}
	fmt.Printf("The reversed string is -%s-\n", string(sl2))
	// variant: using [] int for runes (necessary for Unicode-strings!):
	s := "My Test String!"
	fmt.Println(s, " --> ", reverse(s))
}

func testStrings() {
	s := "Hello World"
	fmt.Println(string(s[:1]))
	chars := []rune(s)
	fmt.Println(string(chars))
	fmt.Println(utf8.RuneCountInString(s))
	new_chars := append(chars, '中')
	fmt.Println(new_chars)
	s = "你好中国"
	fmt.Println(string(s[:1]))
	chars = []rune(s)
	fmt.Println(string(chars[:1]))
	s = "0132"
	fmt.Println(strconv.Atoi(s))
	ints := 133
	s = strconv.Itoa(ints)
	fmt.Println(s[:2])
}
