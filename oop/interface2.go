package main

import "fmt"

// interface{}是万能类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// interface{}如何区分 此时引用的底层数据类型判断
	// 给interface{}提供了一种“类型断言”的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)

		fmt.Printf("value type is %T\n", value)
	}
}

type Books struct {
	auth string
}

func main() {
	book := Books{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc("3.14")

}
