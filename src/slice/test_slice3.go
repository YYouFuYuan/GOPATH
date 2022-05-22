package main

import "fmt"

func main() {
	//声明slice1是切片，并且初始化
	slice1 := []int{1, 2, 3}
	fmt.Printf("len slice1 = %d, slice = %v", len(slice1), slice1)

	//声明但没有分配空间
	var slice2 []int
	//手动分配空间
	slice2 = make([]int, 4) //动态数组，分配空间
	fmt.Printf("len slice2 = %d, slice = %v", len(slice2), slice2)

	if slice1 == nil {
		fmt.Println("slice 空数组")
	} else {
		fmt.Println("不空")
	}

}
