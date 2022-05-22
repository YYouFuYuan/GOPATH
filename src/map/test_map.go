package main

import "fmt"

func main() {
	//第一种 声明map,key是string，vakye是string
	var map1 map[string]string
	//开辟空间，内部也是动态的会自动2倍扩容
	map1 = make(map[string]string, 10)

	map1["one"] = "java"
	map1["two"] = "c++"
	map1["three"] = "go"
	fmt.Println(map1)

	//第二种声明
	map2 := make(map[int]string)
	map2[1] = "java"
	map2[2] = "c++"
	map2[3] = "python"
	fmt.Println(map2)

	//第三种 声明并初始化
	map3 := map[string]string{
		"one":   "java",
		"two":   "java",
		"three": "java",
	}
	fmt.Println(map3)
}
