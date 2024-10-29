package main

import "fmt"

func main() {
	s := []int{1, 2, 3} // len = 3, cap = 3
	// [0, 2)
	s1 := s[0:2] // 里面的元素是[1, 2]
	fmt.Println(s1)

	fmt.Println("=========s1[0]===========")

	// 改变s，s1也会改变
	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)

	// copy可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) // s2 = [0, 0, 0]

	fmt.Println("==============copy=============")
	// 将s中的值，依次拷贝到s2中
	copy(s2, s)
	// 改变s,s2不会改变
	s[0] = 666
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
}
