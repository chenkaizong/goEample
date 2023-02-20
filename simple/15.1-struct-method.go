package main

import "fmt"

type userInfo struct {
	name string
	age  int64
}

//这是一个函数
func getName(userinfo userInfo) string {
	return userinfo.name
}

// 这是一个方法,注意二者的区别，这里也可以给自定义类型添加方法
func (userinfo userInfo) getName() string {
	return userinfo.name
}

// 这个方法是无法修改属性的
func (userinfo userInfo) setName(name string) {
	userinfo.name = name
}

func (userinfo *userInfo) setRealName(name string) {
	userinfo.name = name
}

func main() {
	// 1、结构体是一个值类型，而不是一个引用类型
	a := userInfo{"name", 16}
	b := a
	b.name = "kaka"
	fmt.Printf("%+v\n,%T\n", a, a)
	fmt.Printf("%+v\n,%T\n", b, b)

	fmt.Println(getName(b))
	fmt.Println(b.getName())

	b.setName("白菜")
	fmt.Println(b.getName()) // 修改失败，输出kaka
	b.setRealName("白菜")
	fmt.Println(b.getName()) //修改成功
}
