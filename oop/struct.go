package main

import "fmt"

// 声明一种新的数据类型myint，是int的一个别名
type myint int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	// 构造一个book副本
	book.auth = "jack"
}

func changeBook2(book *Book) {
	// 构造一个book副本
	book.auth = "seal"
}

func main() {
	/*
		var a myint = 10
		fmt.Println("a = ", a)
		fmt.Printf("type of a = %T\n", a)
	*/

	var book1 Book
	book1.title = "Golang"
	book1.auth = "bob"

	fmt.Printf("%v\n", book1)

	changeBook(book1)
	fmt.Printf("%v\n", book1)

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)

}
