package main

import (
	"fmt"
	"time"
)

func goFunc(i int) {
	fmt.Println("goroutine", i, "...")
}

func main() {
	for i := 0; i < 1000; i++ {
		go goFunc(i)
	}
	time.Sleep(time.Second)
}
