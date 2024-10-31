package main

import "fmt"

// 本质是指针
type AnimalIF interface {
	Sleep()
	GetColor() string //  获取颜色
	GetType() string  // 获取种类
}

// 具体类
type Cat struct {
	color string
}

func (c *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (d *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (d *Dog) GetColor() string {
	return d.color
}

func (d *Dog) GetType() string {
	return "Dog"
}

// 接口是一种引用类型，可以直接作为函数参数传递，而不需要使用指针。 所以这里类型<不需要>传 *AnimalIf
func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("Color = ", animal.GetColor())
	fmt.Println("Type = ", animal.GetType())
}

func main() {
	var animal AnimalIF // 接口的数据类型, 父类指针
	animal = &Cat{"Green"}
	animal.Sleep() // 调用的是Cat的Sleep, 多态现象
	animal = &Dog{"Yellow"}
	animal.Sleep()

	cat := Cat{color: "red"}
	dog := Dog{"Gray"}

	showAnimal(&cat)
	showAnimal(&dog)

}
