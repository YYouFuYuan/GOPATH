package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.walk()...")
}

//继承
type SuperMan struct {
	Human //SuperMan继承Human

	level int
}

//重写父类方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func main() {
	human := Human{name: "zhang3", sex: "男"}
	human.Eat()
	human.Walk()

	//定义子类对象
	s := SuperMan{Human{"l14", "男"}, 88}
	s.Walk() //使用父类继承的方法
	s.Eat()  //使用子类的方法，重写了
	s.Fly()  //使用子类自己的方法
}
