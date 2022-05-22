package main

import "fmt"

func myFun(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)
	//interface{} 万能类型

	//万能类型 断言，判断具体是哪个类型
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T\n", value)
	}

}

func main() {
	myFun(100)
	myFun("ay")
}
