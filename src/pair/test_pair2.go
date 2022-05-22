package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	//反射获得实际的类型和值
	fmt.Println("type :", reflect.TypeOf(arg))
	fmt.Println("value :", reflect.ValueOf(arg))
}

type User struct {
	Id   int `info:"id"`
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called...")
}
func reflectObject(arg interface{}) {
	//获取对象的type
	inputType := reflect.TypeOf(arg)
	fmt.Println("input type is ", inputType.Name())

	//获取对象的value
	inputValue := reflect.ValueOf(arg)
	fmt.Println("input value is ", inputValue)

	//获取对象里面的字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)              //得到第i个变量
		value := inputValue.Field(i).Interface() // 得到第i个变量的值
		tag := field.Tag.Get("info")
		fmt.Println("tag = ", tag)
		fmt.Println("field = ", field.Name, " ,type = ", field.Type, " ,value = ", value)
	}

	//获取对象里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s,%v\n", m.Name, m.Type)

	}
}

func main() {
	var num float32 = 1.2
	reflectNum(num)

	user := User{1, "ay", 25}
	reflectObject(user)

}
