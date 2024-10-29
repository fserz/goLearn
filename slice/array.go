package main

import "fmt"

// 定长的数组在传参的时候严格匹配数组类型
func printArray(myArray [4]int) {
	// 仍然是值拷贝
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
	// 值拷贝不会改变原数组
	myArray[0] = 999
}

func main() {

	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}
	fmt.Println("----------------------")
	for index, value := range myArray2 {
		fmt.Println("index = ", index, "value = ", value)
	}

	fmt.Printf("myArray1 Type is = %T\n", myArray1)
	fmt.Printf("myArray2 Type is = %T\n", myArray2)
	fmt.Printf("myArray2 Type is = %T\n", myArray3)

	printArray(myArray3)

	fmt.Println("----------------------")
	for index, value := range myArray3 {
		fmt.Println("index = ", index, "value = ", value)
	}

}
