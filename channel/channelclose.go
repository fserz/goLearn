package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			//close(c) // 报panic错误向关闭的channel发数据
		}
		// close可以关闭一个channel
		close(c) // 不关闭会发生死锁
	}()

	//for {
	//	// ok如果为true便是channel没有关闭，如果为false则已经关闭
	//	if data, ok := <-c; ok {
	//		fmt.Println(data)
	//	} else {
	//		break
	//	}
	//}

	// 可以使用range来迭代不断操作channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished ...")
}
