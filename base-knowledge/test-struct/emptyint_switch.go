package test_struct

import (
	"fmt"
	"reflect"
)

type SpecialString string

func TypeSwitch(whatIsThis any) {
	testFunc := func(any interface{}) {
		// 1. 赋值 any.(type)类似于reelect.Type
		thisAnyType := reflect.TypeOf(any)
		switch thisAnyType.Kind() {
		case reflect.Bool:
			fmt.Printf("any %v is a bool type", reflect.Bool)
		case reflect.Int:
			fmt.Printf("any %v is an int type", reflect.Int)
		case reflect.Float32:
			fmt.Printf("any %v is a float32 type", reflect.Float32)
		case reflect.String:
			fmt.Printf("any %v is a string type", reflect.String)
		case reflect.Struct:
			fmt.Printf("any %v is a special String!", reflect.Struct)
		default:
			fmt.Println("unknown type!")
		}
		fmt.Println("")
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case int:
			fmt.Printf("any %v is an int type", v)
		case float32:
			fmt.Printf("any %v is a float32 type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case SpecialString:
			fmt.Printf("any %v is a special String!", v)
		default:
			fmt.Println("unknown type!")
		}
	}
	testFunc(whatIsThis)
}
