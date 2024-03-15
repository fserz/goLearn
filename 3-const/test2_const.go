package main

import "fmt"

// const定义枚举
const (
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
)

func main() {
	const length int = 10
	fmt.Println("length = ", length)
	//length = 100 常量不允许修改
	fmt.Println("Beijing = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)
	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("k = ", k, "k = ", k)

	// iota只能在const中配合使用
	//var a  int = iota
}
