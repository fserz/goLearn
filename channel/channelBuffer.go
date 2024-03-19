package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
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

	fmt.Println("main 结束")

}
