package main

import (
	"fmt"
	"os"
)

func main() {
	// ...create abort channel...

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	//fmt.Println("Commencing countdown.  Press return to abort.")
	//select {
	//case <-time.After(10 * time.Second):
	//	// Do nothing.
	//case <-abort:
	//	fmt.Println("Launch aborted!")
	//	return
	//}
	//launch()

	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}

}
