package main

import test_struct "base-knowledge/test-struct"

func main() {
	test_struct.TestTag()
	test_struct.TestDayInit()
	var whatIsThis test_struct.SpecialString = "hello world"
	test_struct.TypeSwitch(whatIsThis)

}
