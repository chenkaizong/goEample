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

func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("类型:%v,类型名称:%v,类型种类:%v\n", v, v.Name(), v.Kind())
}

//1、反射获取任意变量的类型
// 2、

func main() {
	a := 10
	b := "100"
	c := true
	d := 23.4

	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

	var e myInt = 20
	var f Person = Person{
		Name: "haha",
		Age:  10,
	}
	reflectFn(e)
	reflectFn(f)

	var g = 25
	reflectFn(&g) //类型名称空

	var h = [3]int{1, 2, 3}
	var i = []int{1, 2, 3}
	reflectFn(h) //类型名称空
	reflectFn(i) //类型名称空
}
