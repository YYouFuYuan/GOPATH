package main

import "fmt"

//接口
type Reader interface {
	ReadBook()
}

//接口
type Writer interface {
	WriteBook()
}

//具体类
type Book struct {
}

//实现两个接口
func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	//b:pair<type:*Book,value:Book的地址>
	b := &Book{}
	fmt.Printf("b type %T\n", b)

	//接口1
	var r Reader
	r = b //r:pair<type:*Book,value:Book地址> 父类指针指向子类对象,但对象的本质类型是没有发生改变的
	r.ReadBook()

	//接口2
	var w Writer
	w = r.(Writer) //r本质的类型是 *Book，所以可以断言为Writer，类似C++的强转，向上强转
	w.WriteBook()
}
