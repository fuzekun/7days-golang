package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	testReflect()
	myPrint(Day(1), "was", Celsius(18.36)) // Tuesday was 18.4 °C
}

func testReflect() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// setting a value:
	// v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of v:", v.Type(), reflect.TypeOf(x).Kind())
	fmt.Println("settability of v:", v.CanSet())
	// 使用relect修改值的核心点，对v.Elem().setFloat()
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	if reflect.TypeOf(x).Kind() == reflect.Float64 {
		// 类型的Kind是float类型，所以说是生效的
		v.SetFloat(3.1415) // this works!
	} else if v.Type().Kind() == reflect.Float64 {
		// 由于值的类型是float* 类型，所以说这个是不生效的
		v.SetFloat(3.146)
	}
	fmt.Println(v.Interface())
	fmt.Println(v)
}

type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + " °C"
}

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func myPrint(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			os.Stdout.WriteString(" ")
		}
		switch a := arg.(type) { // type switch
		case Stringer:
			os.Stdout.WriteString(a.String())
		case int:
			os.Stdout.WriteString(strconv.Itoa(a))
		case string:
			os.Stdout.WriteString(a)
		// more types
		default:
			os.Stdout.WriteString("???")
		}
	}
}
