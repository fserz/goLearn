package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
		// 终止当前的goroutine
		//runtime.Goexit()
	}
}

// 主goroutine
func main() {
	// 创建一个go进程 去执行newTask()
	go newTask()

	//fmt.Println("main goroutine exit")

	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}

}
