package main

import "fmt"

func main() {
	// 声明slice1是切片，初始化为1,2,3，len=3
	//slice1 := []int{1, 2, 3}

	// 声明slice1是切片，但是没有为slice分配空间
	//var slice1 []int
	//slice1 = make([]int, 3)
	//slice1[0] = 100

	//var slice1 []int = make([]int, 3)

	// := 推导出 slice1是切片
	slice1 := make([]int, 3)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 判断slice是否为空
	if slice1 == nil {
		fmt.Println("slice1是空切片的")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}
