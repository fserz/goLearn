package main

import (
	"fmt"
)

func recoverFunc() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func someFunc() {
	defer recoverFunc()
	fmt.Println("Start of someFunc")
	panic("something went wrong")
	fmt.Println("End of someFunc") // 这行代码不会执行，因为panic导致函数中止
}

func main() {
	fmt.Println("Start of main")
	someFunc()
	fmt.Println("End of main")
}
