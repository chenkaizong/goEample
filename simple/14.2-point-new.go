package main

import "fmt"

func main() {
	// 下面这三行的写法是错误的,因为没有分配存储空间
	// var a *int
	// *a = 10
	// fmt.Println(*a)

	// 通过new 关键词来给a 分配存储空间
	var a = new(int)
	*a = 10
	fmt.Println(*a)

}
