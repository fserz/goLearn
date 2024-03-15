package main

// 起别名
//import aa "fmt"
// 无法使用当前包方法，但会执行init
//import _ "fmt"
// .是指讲所有方法导入至本包，一般不要使用，容易同名歧义
import . "fmt"

func main() {
	//lib1.Lib1Test()
	//lib2.Lib2Test()

	//aa.Println(111)
	Println(222)
}
