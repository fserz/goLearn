package main

import "fmt"

// map 声明方式
func main() {
	// 第一种方式
	// 声明myMap1是一种map，key是string，value也是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空的map")
	}

	// 使用map前，需要先使用make给map分配数据空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "C++"
	myMap1["three"] = "py"

	fmt.Println(myMap1)

	// 第二种方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "go"
	myMap2[3] = "C++"
	fmt.Println(myMap2)

	// 第三种方式
	myMap3 := map[string]string{
		"one":   "go",
		"two":   "java",
		"three": "go",
	}
	println(myMap3)

}
