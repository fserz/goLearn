package main

import (
	"fmt"
	"time"
)

func main() {
	// 有缓存channel
	c := make(chan int, 3)
	// 查看当前channel容量cap，元素数量len
	fmt.Println("channel初始化 len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		// 子 goroutine 发送第四个值 3 时，由于 channel 已满，发送会阻塞，直到主 goroutine 从 channel 中接收一个值。
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go进程正在运行：len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c // 从c中接受数据，并赋值给num
		fmt.Println("num = ", num)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("main 结束")

}
