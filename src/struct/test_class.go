package main

import "fmt"

//方法 变量和类名，大小写都是有差别的，大写外部可以访问，小写是不行的
type Hero struct {
	Name string
	Ad   int
}

//定义类的方法，这里Hero要加*，否则this只是一个副本，无法真正改变对象的值
func (this *Hero) Show() {
	fmt.Printf("%v\n", this)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(name string) {
	this.Name = name
}

func main() {
	//创建对象
	hero := Hero{Name: "Ay", Ad: 100}
	hero.Show()
	hero.SetName("ay")
	hero.Show()
}
