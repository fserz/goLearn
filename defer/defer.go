package main

import "fmt"

func main() {
	// 结束之前的一个机制
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main: hello go 1")
	fmt.Println("main: hello go 2")
}
