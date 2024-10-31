package main

import (
	"fmt"
	"reflect"
)

// reflect 机制
func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("type: ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345
	str := "test string"
	num1 := 666

	reflectNum(num)
	reflectNum(num1)
	reflectNum(str)
}
