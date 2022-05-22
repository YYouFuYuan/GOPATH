package main

import "fmt"

func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

func swapP(a *int, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}
func main() {
	var a int = 10
	var b int = 20
	//值传递
	swap(a, b)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//指针传递
	swapP(&a, &b)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//一级指针
	var p *int = &a
	fmt.Println("p = ", p)
	fmt.Println("&a = ", &a)

	//二级指针
	var pp **int = &p
	fmt.Println("pp = ", pp)
	fmt.Println("&p = ", &p)
	fmt.Println("**pp = ", **pp)
	fmt.Println("*p = ", *p)
	fmt.Println("a = ", a)
}
