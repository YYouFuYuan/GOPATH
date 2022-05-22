package main

import "fmt"

/**
动态数组 传递指针
*/
func printArray2(array []int) {
	//匿名参数，跟python一样
	for _, value := range array {
		fmt.Println("value = ", value)
	}
	//修改值,引用传递。传递指针
	array[0] = 100
}
func main() {
	//动态数组
	myArray := []int{1, 2, 3, 4}
	printArray2(myArray)
	fmt.Println(myArray[0])

}
