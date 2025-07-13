package main

import (
	"base-knowledge/variable"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"
)

func main() {
	testInt()
	testTime()
	testStrings()

}

func testInt() {
	// 1. 如果是32范围的整数，直接加，也不会溢出。因为平台64位，并且默认赋值会推断成int类型的。也就是64位
	a := math.MaxInt32
	a++
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	// 2. 如果显示的指定int32类型的整数，就会溢出了
	var b int32 = math.MaxInt32
	b++
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))

	// 3. 需要注意的是，如果有用到32位的平台，最好指定int32，或者int64，这样才不会有溢出的风险
}
func testTime() {
	var week time.Duration
	t := time.Now()
	fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// 21.12.2011
	t = time.Now().UTC()
	fmt.Println(t)          // Wed Dec 21 08:52:14 +0000 UTC 2011
	fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC))  // Wed Dec 21 08:56:34 2011
	// The time must be 2006-01-02 15:04:05
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
}

func testStrings() {
	// 1. 字符串分割
	s := "hello world"
	splits := strings.Split(s, " ")
	fmt.Println(splits)
	for _, split := range splits {
		fmt.Println(split)
	}
	// 2. 字符串拼拼接
	// 拼接大量日志
	var builder strings.Builder
	builder.Grow(1024) // 预分配1KB
	builder.WriteString("[INFO] ")
	builder.WriteString(time.Now().Format("2006-01-02"))
	builder.WriteString(" Operation completed")
	fmt.Println(builder.String())

	// 3. 基于strings.builder实现线程安全的stringbuffer类
	buffer := variable.StringBuffer{}
	buffer.WriteString("[INFO] ")
	buffer.WriteString(time.Now().Format("2006-01-02"))
	buffer.WriteString(" Operation completed")
	fmt.Println(buffer.String())
}
