package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"` //标签，类似注解
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"xing", "zang"}}

	//结构体 => json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s \n", jsonStr)

	//json => 结构体
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Println(my_movie)
}
