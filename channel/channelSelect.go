package main

import (
	"fmt"
	"time"
)

// c, quit都是chan类型的简写
func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			// 如果c可写，则case就会进来
			x = y
			y = x + y
		case <-quit:
			time.Sleep(5 * time.Second)
			fmt.Println("quit...")
			return
		}
	}
}

// sub go
func Run(c, quit chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	// Run执行完毕后向quit发值，表示任务完成，进入case <-quit
	quit <- 0
}

//程序的运行流程是，同时启动了一个 goroutine：从通道 c 中读取数据并打印。
//main 函数中调用的 fibonacci(c, quit) ，在主 goroutine 中执行的，用于生成斐波那契数列并向通道 c 中发送数据

func main() {
	c := make(chan int)
	quit := make(chan int)

	// 启动一个匿名goroutinue
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()

	go Run(c, quit)

	// main go
	fibonacii(c, quit)

}
