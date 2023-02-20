package main

import "fmt"

// 1、type 自定义类型，输出类型的值是 myint
type myint int64

// 2、类型别名,别名用等于号，输出类型的时候是 原来的类型

type myfloat = float64

// 3、结构体

type userInfo struct {
	name     string
	age      int64
	sex      string
	UserName string
}

func main() {
	var a myint = 100
	fmt.Printf("%d,%T,%v,\n", a, a, a)
	var b myfloat = 100
	fmt.Printf("%f,%T,%v,\n", b, b, b)
	// 方式1
	var c userInfo = userInfo{
		name: "haha",
		age:  16,
		sex:  "男",
	}
	c.name = "白菜"
	fmt.Printf("%v\n,%+v,%T\n", c, c, c)
	// 方式2
	var d = new(userInfo)
	d.name = "bb"
	d.age = 20
	d.sex = "男"
	fmt.Printf("%v\n,%+v,%T\n", d, d, d)
	// 方式3
	var f = &userInfo{}

	f.name = "bb"
	f.age = 20
	f.sex = "男"
	fmt.Printf("%v\n,%+v,%T\n", f, *f, f)

	// 方式4

	g := userInfo{"mingzi", 22, "女", "LAL"}
	fmt.Printf("%v\n,%+v,%T\n", g, g, g)

}
