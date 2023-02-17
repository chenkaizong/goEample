package main

import "fmt"

type calc func(int, int) int

type myInt int //自定义类型

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func computer(c calc, x int, y int) int {
	return c(x, y)
}

func operator(operat string) calc {
	var calcFun calc
	switch operat {
	case "+":
		calcFun = add
	case "-":
		calcFun = sub
	default:
		calcFun = sub
	}
	return calcFun
}

func main() {
	// 1、函数作为变量
	var c calc
	c = add
	fmt.Printf("%T\n", c)

	fmt.Println(c(3, 5))
	// 2、函数作为参数

	fmt.Println(computer(c, 3, 5))

	//3、匿名函数
	fmt.Println(computer(func(i1, i2 int) int {
		return i1 * i2
	}, 3, 5))

	//4、返回一个函数

	fmt.Println(operator("+")(3, 4))
}
