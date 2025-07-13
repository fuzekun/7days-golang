package main

import (
	"base-knowledge/function"
	"base-knowledge/variable"
	"fmt"
)

func main() {
	// 测试函数闭包
	func() { fmt.Println("hello world") }()
	// 测试记忆化搜索
	function.TestFibonacci()
	// 测试指针引用和值引用
	testPoint()
	// 测试不定长变量
	testVariablesLenNotFixed()
}

func testVariablesLenNotFixed() {
	f1("付泽坤", "赋能", "go语言开发")
}

// 测试指针引用和值引用
func testPoint() {
	buffer := variable.StringBuffer{}
	bufferPointer := new(variable.StringBuffer)
	testCall(buffer, bufferPointer)
	fmt.Println("值引用结果:", buffer.String())
	fmt.Println("指针引用结果:", (*bufferPointer).String())
}

func testCall(buffer variable.StringBuffer, bufferPointer *variable.StringBuffer) {
	bufferPointer.WriteString("hello world")
	buffer.WriteString("hello world")
}

func f1(s ...string) {
	f2(s...)
	f3(s)
}

func f2(s ...string) {
	fmt.Println(s)
}
func f3(s []string) {}
