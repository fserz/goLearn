package main

import (
	"fmt"
	"time"
)

//	select具备多路channel的监控功能
//		select {
//		case <- chan1:
//			// 如果chan1成功读到数据，则进行该case处理语句
//		case chan2 <- 1:
//			// 如果向chan2写入数据，则进行该case处理语句
//		default
//			// 如果上面都没有成功，则进入default处理流程
//		}

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

	// 这两行**不能颠倒**。没有缓冲的情况下，发送和接收必须配对。
	//如果先执行 fibonacii，而 Run 还没有开始接收数据，就会导致程序阻塞，甚至可能出现死锁。
	// 如果希望颠倒顺序，可以将 c 改为有缓冲的 channel
	go Run(c, quit)
	// main go
	fibonacii(c, quit)

}
