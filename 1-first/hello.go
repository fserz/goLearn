package main // 程序包名

import (
	"context"
	"fmt"
	"time"
)

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")
	// context 使用场景
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, "node01")
	go worker(ctx, "node02")
	go worker(ctx, "node03")
	go worker(ctx, "node04")
	time.Sleep(10 * time.Second)
	fmt.Println("stop the goroutine")
	// 调用cancel函数触发ctx.Done()
	cancel()
	time.Sleep(5 * time.Second)

}

func worker(ctx context.Context, name string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "got the stop channel")
				return
			default:
				fmt.Println(name, "still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
}
