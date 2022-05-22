package main

import "fmt"

/**
固定数组传参只能传固定长度,并且是值拷贝
*/
func printArray(array [10]int) {
	for index, value := range array {
		fmt.Println("index = ", index, ",value = ", value)
	}
	//array和外面的myArray2是不同的东西了，已经不是同一个地址了，不改变原本的值
	array[0] = 100
}
func main() {
	//固定长度的数组
	var myArray1 [10]int
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	//初始化,未初始的默认为0
	myArray2 := [10]int{1, 2, 3, 4, 5}
	for index, value := range myArray2 {
		fmt.Println("index = ", index, ",value = ", value)
	}
	printArray(myArray2)
	fmt.Println(myArray2[0])

}
