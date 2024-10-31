package main

import "fmt"

// 如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	// 如果类的属性首字母大写，表示属性可以对外访问，否则的话只能够在类的内部访问
	Name  string
	Ad    int
	Level int
}

//	func (this Hero) Show() {
//		fmt.Println("Name = ", this.Name)
//		fmt.Println("Ad = ", this.Ad)
//		fmt.Println("Level = ", this.Level)
//	}
//
//	func (this Hero) GetName() string {
//		fmt.Println("name = ", this.Name)
//		return this.Name
//	}
//func (this Hero) SetName(newName string) {
//	// this是调用该方法的对象的一个副本（拷贝）
//	this.Name = newName
//}

// 在 Go语言中，接收者的名称用于在方法主体中引用该接收者的实例。
// 可以用h, this各种标识均可
func (h *Hero) Show() {
	fmt.Println("Name = ", h.Name)
	fmt.Println("Ad = ", h.Ad)
	fmt.Println("Level = ", h.Level)
}
func (h *Hero) GetName() string {
	fmt.Println("name = ", h.Name)
	return h.Name
}

// 指针传递可以修改对象中的值
func (h *Hero) SetName(newName string) {
	h.Name = newName
}

func main() {
	hero := Hero{Name: "zhangsan", Ad: 100, Level: 1}
	hero.Show()
	hero.SetName("li4")
	hero.Show()
}
