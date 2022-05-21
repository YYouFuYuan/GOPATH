package main

import "fmt"

//定义枚举作用
const (
	//iota  默认值为0， 每一行的iota会自动累加, iota只能配合const使用
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = 1, b = 2
	c, d                      // iota = 1, c = 2, d = 3
)

func main() {
	fmt.Println("hello,go")
	const length = 10
	fmt.Println("length = ", length)
	fmt.Println("shenzhen = ", SHENZHEN)
	fmt.Printf("c = %d,d = %d", c, d)
}
