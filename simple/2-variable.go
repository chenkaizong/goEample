package main

import "fmt"

func test() (name string, age int) {
	return "NAME", 20
}

func main() {
	// 1 可以全局定义变量的方式 var
	//1.1 var 先声明再赋值
	var a string
	a = "哈哈"
	fmt.Println(a)
	// 1.2 var 声明后直接赋值
	var b string = "啦啦"
	fmt.Println(b)
	// 1.3 一次声明多个相同类型
	var c1, c2 string = "C1", "C2"
	c1 = "c1"
	c2 = "c2"
	fmt.Println(c1, c2)
	// 1.4 一次声明多个不同类型
	var (
		d1 int    = 20
		d2 string = "哈哈"
		d3 bool   = false
	)
	fmt.Println(d1, d2, d3)

	// 2 短变量方式，这只能在局部定义，他可以自己判断数据类型

	e := 100
	fmt.Println(e)
	fmt.Printf("%T\n", e)

	f1, f2 := "我去", "蛋碎"
	fmt.Println(f1, f2)

	//3、匿名变量，函数返回多个值时候使用匿名变量,必须使用_ 来让程序忽略多余的返回值
	name, _ := test()
	fmt.Println(name)

}
