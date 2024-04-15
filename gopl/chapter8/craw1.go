package main

import (
	"fmt"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//func main() {
//	worklist := make(chan []string)
//
//	// Start with the command-line arguments.
//	go func() { worklist <- os.Args[1:] }()
//
//	// Crawl the web concurrently.
//	seen := make(map[string]bool)
//	for list := range worklist {
//		for _, link := range list {
//			if !seen[link] {
//				seen[link] = true
//				go func(link string) {
//					worklist <- crawl(link)
//				}(link)
//			}
//		}
//	}
//}

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
