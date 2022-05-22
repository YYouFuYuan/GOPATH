package main

import "fmt"

//定义接口  本质是指针
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物颜色
	GetType() string  //获取动物类型
}

//具体的类 不用显示的语法去继承
type Cat struct {
	color string //猫的颜色
}

//实现接口3个方法
func (this *Cat) Sleep() {
	fmt.Println("cat i sleep()")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

//具体的类 不用显示的语法去继承
type Dog struct {
	color string //猫的颜色
}

//实现接口3个方法
func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep()")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func main() {

	var animal AnimalIF    //父类指针
	animal = &Cat{"Green"} //指向子类对象

	animal.Sleep() //调用cat

	dog := Dog{"white"}
	animal = &dog
	animal.Sleep() //调用dog
}
