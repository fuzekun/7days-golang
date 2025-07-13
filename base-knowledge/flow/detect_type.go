package flow

import (
	"reflect"
	"time"
)

// GetType 识别变量的基本类型，返回类型名称字符串
func GetType(v interface{}) string {
	if v == nil {
		return "nil"
	}

	// 获取反射值和类型
	rv := reflect.ValueOf(v)
	rt := rv.Type()

	// 解引用指针
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return rt.String() + " (nil pointer)"
		}
		rv = rv.Elem()
		rt = rv.Type()
	}

	// 判断基本类型
	switch rv.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "int (" + rt.String() + ")"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return "uint (" + rt.String() + ")"
	case reflect.Float32, reflect.Float64:
		return "float (" + rt.String() + ")"
	case reflect.Complex64, reflect.Complex128:
		return "complex (" + rt.String() + ")"
	case reflect.String:
		return "string"
	case reflect.Array:
		return "array (" + rt.String() + ")"
	case reflect.Slice:
		return "slice (" + rt.String() + ")"
	case reflect.Map:
		return "map (" + rt.String() + ")"
	case reflect.Struct:
		// 特殊处理时间类型
		if _, ok := v.(time.Time); ok {
			return "time.Time"
		}
		return "struct (" + rt.String() + ")"
	case reflect.Interface:
		return "interface (" + rt.String() + ")"
	case reflect.Func:
		return "function (" + rt.String() + ")"
	case reflect.Chan:
		return "channel (" + rt.String() + ")"
	case reflect.Ptr:
		return "pointer (" + rt.String() + ")"
	default:
		return "unknown (" + rt.String() + ")"
	}
}
