package main

import "fmt"

// map的使用方式
func printMap(cityMap map[string]string) {
	// cityMap是一个引用传递
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	// delete
	delete(cityMap, "China")

	fmt.Println("===================")

	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	// 修改
	cityMap["USA"] = "DC"
	fmt.Println("===================")

	//遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	fmt.Println("===================")
	printMap(cityMap)

	fmt.Println("===================")
	changeValue(cityMap)
	printMap(cityMap)

}
