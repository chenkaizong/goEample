package main

import "fmt"

// 空接口表示任意类型
type Empty interface {
}
type Phone struct {
	Name string
}

func (phone Phone) start() {
	fmt.Println("start")
}

func show(i interface{}) {
	fmt.Printf("%+v\n", i)
}

func main() {

	//1、空接口可以赋值给任意类型
	var p1 Empty
	var str string = "字符串"
	p1 = str

	fmt.Printf("值%v 类型%T\n", p1, p1)

	//2、空接口可以直接定义
	var p2 interface{}
	p2 = false

	fmt.Printf("值%v 类型%T\n", p2, p2)

	//3.1、功能：用于函数输入输出任意类型

	show(1)
	show(map[string]string{"key": "value"})

	//3.2 功能 用于定义任意切片或map
	var p3 = make(map[string]interface{})
	var p4 = []interface{}{"100", false}
	p3 = map[string]interface{}{
		"key":    "value",
		"number": 1,
	}
	show(p3)
	show(p4)

	//4、断言 1 判断是否是string类型，并返回结果
	v, err := p1.(string)
	println(v, err) //字符串 true

	// 4.2 断言2：type-switch

	switch x := p1.(type) {
	case string:
		fmt.Println(x, "是字符串")

	}

}
