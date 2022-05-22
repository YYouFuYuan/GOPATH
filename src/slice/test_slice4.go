package main

import "fmt"

func main() {
	//一开始申请5个空间，只初始化前3个空间，len是实际数组长度，cap是整个数组容量
	var numbers []int = make([]int, 3, 5)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)

	//动态添加元素
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)

	//cap如果已经满了，此时会进行双倍扩容
	numbers = append(numbers, 3)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)

	//切片截取
	s := numbers[3:6] //此时的s和number是指向同一空间的，改变s也会改变number
	fmt.Println(s)

	//深拷贝截取
	s2 := make([]int, 3)
	copy(s2, s)
	s2[0] = 100
	fmt.Println("numbers = ", numbers)
	fmt.Println("s = ", s)
	fmt.Println("s2 = ", s2)
}
