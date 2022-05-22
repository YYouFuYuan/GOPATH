package main

import "fmt"

//给类型起别名
type myint int

// 定义结构体
type Book struct {
	title string
	auth  string
}

//结构体传参时，是传值，不是指针,改变内部不影响外部
func changeBook(book Book) {
	book.title = "666"
}

func changeBook2(book *Book) {
	book.title = "666"
}
func main() {
	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"
	fmt.Println("%v\n", book1)
	changeBook(book1)
	fmt.Println("%v\n", book1)

	changeBook2(&book1)
	fmt.Println("%v\n", book1)
}
