package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string, 4)
	//out := make(chan string)
	ch <- "A"
	fmt.Println("len(ch) = ", len(ch))
	ch <- "B"
	ch <- "C"
	ch <- "D"
	fmt.Println("cap(ch) = ", cap(ch))
	fmt.Println(<-ch)
	fmt.Println("len(ch) = ", len(ch))
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("len(ch) = ", len(ch))
	fmt.Println(<-ch)
	fmt.Println("cap(ch) = ", cap(ch))
	fmt.Println("len(ch) = ", len(ch))
	res := mirroredQuery()
	fmt.Println(res)

}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses // return the quickest response
}

func request(hostname string) (response string) { /* ... */
	fmt.Println(response)
	return "This is a response"
	sync.WaitGroup{}
}
