package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human // SuperMan类继承了Human类的方法
	level int
}

// 重定义父类的方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan Eat....")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main() {
	h := Human{"zhangsan", "female"}
	h.Walk()
	h.Eat()

	// 定义一个子类对象
	//s := SuperMan{Human{"wang5", "male"}, 2}

	var s SuperMan
	s.name = "Tom"
	s.sex = "male"
	s.level = 88

	s.Walk()
	s.Eat()
	s.Fly()

	s.Print()
}
