package test_struct

import (
	"fmt"
	"reflect"
)

type TagType struct { // tags
	field1 bool   "这个是field1的注解"
	field2 string "这个是field2的注解"
	field3 int    "这个是field3的注解"
}

func TestTag() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	// 1. 反射：在运行的时候，可以获取数据的元信息，也就是数据的类型信息
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Println(ixField.Tag, ixField.Name, ixField.Type, ixField.Type.Kind())
}
