package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}
type myInt int

func reflectExample(x interface{}) {
	// fmt.Println(x + 10) 报错，x不是int
	// 通过断言
	var i, err = x.(int)
	if err != false {
		fmt.Println(i + 10)
	}

	// 通过反射
	v := reflect.ValueOf(x)
	fmt.Println(v.Int() + 11)
}

func reflectFn(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Int:
		fmt.Println(v.Int() + 10)
	case reflect.Float64:
		fmt.Println(v.Float() + 11.2)
	case reflect.String:
		fmt.Println(v.String() + "hhh")
	}
}

func main() {
	a := 10
	reflectExample(a)

	var b int = 100
	var c float64 = 100.1
	var f string = "11111"
	reflectFn(b)
	reflectFn(c)
	reflectFn(f)

}
