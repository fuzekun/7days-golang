package main

import (
	"base-knowledge/flow"
	"base-knowledge/variable"
	"fmt"
)

func main() {
	buffer := variable.StringBuffer{}
	buffer.WriteString("hello world")
	bufferType := flow.GetType(buffer)
	fmt.Println("buffer类型:", bufferType)

}
