package main

import "fmt"

func fool(a string, b int) int {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := 254

	return c
}

// 返回多个值，匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 4344, 232
}

// 返回多个值，有形参
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----foo3----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 100
	r2 = 200
	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("----foo4----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// r1,r2作用域空间是foo3整个函数体的{}空间
	fmt.Println("r1  = ", r1)
	fmt.Println("r2  = ", r2)

	r1 = 444
	r2 = 555
	return
}

func main() {
	c := fool("abc", 299)
	fmt.Println(c)
	ret1, ret2 := foo2("hhahah", 8909)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)
	ret3, ret4 := foo3("xxxx", 1004)
	fmt.Println("ret3 = ", ret3, "ret4", ret4)
	ret5, ret6 := foo4("1sd2asd3", 234)
	fmt.Println("ret5 = ", ret5, "ret6 = ", ret6)
}
