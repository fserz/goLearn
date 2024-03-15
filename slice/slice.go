package main

import "fmt"

func printArray1(myArray []int) {
	// _表示匿名变量
	for _, value := range myArray {
		fmt.Println("value =", value)
	}
	myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4} // 动态数组，切片

	fmt.Printf("myArray type is %T\n", myArray)

	printArray1(myArray)

	fmt.Println("=============================")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

}
