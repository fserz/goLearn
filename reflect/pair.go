package main

import "fmt"

func main() {
	var a string
	// pair<static type:string, value:"abc>
	a = "abc"

	// pair<type:string, value:"abc>
	var allType interface{}
	allType = a

	str, isStr := allType.(string)
	fmt.Println(str)
	fmt.Println(isStr)

}
