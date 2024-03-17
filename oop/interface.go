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

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func shwoAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("Color = ", animal.GetColor())
	fmt.Println("Type = ", animal.GetType())
}

func main() {
	var animal AnimalIF // 接口的数据类型
	animal = &Cat{"Green"}
	animal.Sleep() // 调用的是Cat的Sleep
	animal = &Dog{"Yellow"}
	animal.Sleep()

	cat := Cat{color: "red"}
	dog := Dog{"Gray"}

	shwoAnimal(&cat)
	shwoAnimal(&dog)

}
