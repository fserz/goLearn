package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (h *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (h *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human // SuperMan类继承了Human类的方法
	level int
}

// 重定义父类的方法
func (s *SuperMan) Eat() {
	fmt.Println("SuperMan Eat....")
}

// 不重写则调用父类方法
//func (s *SuperMan) Walk() {
//	fmt.Println("SuperMan Walk fast....")
//}

func (s *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (s *SuperMan) Print() {
	fmt.Println("name = ", s.name)
	fmt.Println("sex = ", s.sex)
	fmt.Println("level = ", s.level)
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
