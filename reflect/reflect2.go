package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called")
	fmt.Printf("%v\n", this)
}

func DoFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	//  通过type获得里面的字段
	// 1. 获取interface的reflect.Type,通过Type传到NumFiled
	// 2. 得到每个field，数据类型
	// 3. 通过filed有一个interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		filed := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", filed.Name, filed.Type, value)

	}

	// 通过Type获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Println("%s : %v\n", m.Name, m.Type)
	}
}

func main() {
	user := User{1, "Aceld", 18}
	DoFiledAndMethod(user)

}
