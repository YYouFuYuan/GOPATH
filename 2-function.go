package main

import "fmt"

/**
单一返回值
*/
func fool(name string, age int) int {
	fmt.Println("name = ", name)
	fmt.Println("age = ", age)
	c := 100
	return c

}

/**
返回多个返回值,且返回值使用名称
*/
func foo2(a string, b int) (r1 int, r2 string) {
	r1 = 20
	r2 = "Ay"
	return

}
func main() {
	c := fool("ay", 25)
	fmt.Println("c = ", c)

	age, name := foo2("ay", 30)
	fmt.Println("age:", age)
	fmt.Println("name:", name)
}
