package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// 给interface{}提供了一种类型断言的机制
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
