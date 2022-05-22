package main

import "fmt"

func main() {
	//方法一：声明变量，默认值是0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of a=%T\n", a) //打印类型 printf 格式化输出

	//方法二：声明变量并赋值
	var b int = 100
	fmt.Println("b=", b)

	//方法三：初始化时，省去数据类型，通过值自动匹配
	var c = 100
	fmt.Println("c=", c)

	//方法四：省去va关键字，类型自动匹配（最常用）,但不能用于全局变量声明
	e := 100.0
	fmt.Printf("type of e %T\n", e)
}
