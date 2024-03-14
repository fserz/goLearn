package main

import "fmt"

var va int = 100
var vb = 200

// :无法全局
//vc := 300

func main() {
	// 变量默认值为0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	// 声明变量
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	// 不推荐
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	// 最常用的,省略var，只能用在函数体内
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("va = ", va)
	fmt.Println("va = ", vb)
	//fmt.Println("vc = ", vc)

	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)
	var kk, ll = 100, "Alice"
	fmt.Println("kk = ", kk, "ll = ", ll)

	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv)
	fmt.Println("jj = ", jj)

}
