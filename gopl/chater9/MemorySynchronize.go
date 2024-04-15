package main

import (
	"fmt"
	"time"
)

//两个goroutine是并发执行，并且访问共享变量时也没有互斥，
//会有数据竞争，所以程序的运行结果没法预测的话也请不要惊讶。

func main() {
	var x, y int
	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()

	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()

	time.Sleep(time.Second)
}
