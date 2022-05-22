package main

import "fmt"

func printMap(cityMap map[string]string) {
	//遍历,传递的是指针，指向同一个地址
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func main() {
	cityMap := make(map[string]string)
	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	printMap(cityMap)

	//删除
	delete(cityMap, "China")

	//修改
	cityMap["USA"] = "DC"

	printMap(cityMap)
}
