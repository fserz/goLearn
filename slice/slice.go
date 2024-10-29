package main

import "fmt"

// 动态数组在传参时是引用传递，而且不同长度的动态数组他们的形参是一致的。
func printArray1(myArray []int) {
	// _表示匿名变量
	for _, value := range myArray {
		fmt.Println("value =", value)
	}
	myArray[0] = 666
}

func main() {
	myArray := []int{1, 2, 3, 4} // 动态数组，切片 slice

	fmt.Printf("myArray type is %T\n", myArray)

	printArray1(myArray)

	fmt.Println("=============================")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

}
