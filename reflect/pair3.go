package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book2 struct {
}

func (this *Book2) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book2) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {

	// b: pair<type:Book, value:book{}地址>
	b := &Book2{}

	// r: pair<type:Book, value:>
	var r Reader

	// r: pair<type:Book, value:book{}地址>
	r = b

	r.ReadBook()

	var w Writer

	// r: pair<type:Book, value:>
	w = r.(Writer) // 此处的断言为什么成功？因为w r具体的type是一致的

	w.WriteBook()
}
