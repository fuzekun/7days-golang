package main

import (
	"errors"
	"fmt"
	"strings"
)

// java中的函数式接口

type Any interface{}
type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
	// ...
}

func (c *Car) String() string {
	return fmt.Sprintf("{Model: %s, Manu: %s, Year: %d}", c.Model, c.Manufacturer, c.BuildYear)
}

type Cars []*Car

func main() {
	// make some cars:
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMWs", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}
	// query:
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allCarsPointer := &allCars
	xlModelCar := allCarsPointer.FindAll(func(car *Car) bool {
		return (strings.Contains(car.Model, "XL"))
	})
	fmt.Println("XLModelCar:", xlModelCar)
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)
	manufacturesAny := allCarsPointer.Map(func(car *Car) Any { return car.Manufacturer })
	manufacturers, _ := ChangeAnyArrayToRealArray[string](manufacturesAny)
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")

	cars := []Any{
		&Car{Model: "Fiesta", BuildYear: 2008},
		&Car{Model: "X5", BuildYear: 2022},
	}

	// 通用的映射方法，将某一个类，映射成
	models, err := MapArray[string](cars, func(c Any) Any {
		return c.(*Car).Model
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(models)

	nums := []int{2, 2}
	testArray(nums)
	fmt.Println(nums)
}

// 封装遍历的过程，省略了一个for循环而已
// Process all cars with the given function f:
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

// Find all cars matching a given criteria.
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	// 为什么这里还需要包一层func，实际上Process省略了外层for循环（遍历），里面不做任何处理。
	// 里面的处理逻辑还是需要自己设计的，所以只要有逻辑，就需要自己定义，然后调用它
	// 这里实际上f并不是处理逻辑，而只是处理逻辑的一部分。真正的逻辑是 if f(c) {do something} ，所以需要重新定义一个func去包裹这个逻辑
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

// 处理映射根据某一个字段映射的逻辑
// Process cars and create new data.
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, 0)
	cs.Process(func(c *Car) {
		result = append(result, f(c))
	})
	return result
}

// 返回一个Appender，和钩子结果，当使用Processor处理了这个方法之后。钩子结果就会执行了。起到一个延迟执行的效果
func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	// Prepare maps of sorted cars.
	sortedCars := make(map[string]Cars)

	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	// Prepare appender function:
	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}

func MakeSortedAppenderNew(manufacturers []string, cars *Cars) map[string]Cars {
	// Prepare maps of sorted cars.
	sortedCars := make(map[string]Cars)

	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	// 进行方法的构造
	cars.Process(func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	})

	return sortedCars
}

// 通用的映射方法，将某一个类，映射成对应的字段数组
func MapArray[T Any](array []Any, f func(c Any) Any) ([]T, error) {
	// 1. 首先进行映射
	n := len(array)
	result := make([]Any, n)
	for i, v := range array {
		result[i] = f(v)
	}
	// 2. 将映射后的数组，转换成给定的类型
	return ChangeAnyArrayToRealArray[T](result)
}

func ChangeAnyArrayToRealArray[T any](array []Any) ([]T, error) {
	if len(array) == 0 {
		return make([]T, 0), nil
	}
	result := make([]T, len(array))
	for i, v := range array {
		if val, ok := v.(T); ok {
			result[i] = val
		} else {
			return nil, errors.New("invalid type")
		}
	}
	return result, nil
}

func testArray(nums []int) {
	nums[0] = 1
}

/* Output:
AllCars:  [0xf8400038a0 0xf840003bd0 0xf840003ba0 0xf840003b70]
New BMWs:  [0xf840003bd0]
Map sortedCars:  map[Default:[0xf840003ba0] Jaguar:[] Land Rover:[] BMW:[0xf840003bd0 0xf840003b70] Aston Martin:[] Ford:[0xf8400038a0]]
We have  2  BMWs
*/
