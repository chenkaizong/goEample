package main

import "fmt"

// 1、默认方式
func sumFn(x int, y int) int {
	return x + y
}

// 1.1参数类型简写，只能简写前面的
func subFn(x, y int) int {
	return x - y
}

// 2.1 可变参数
func sumFn1(x ...int) int {
	var result int
	for _, v := range x {
		result += v
	}
	return result
}

// 2.1 可变参数
func sumFn2(x int, y ...int) int {
	var result int = x
	for _, v := range y {
		result += v
	}
	return result
}

// 3.1 返回多个值
func calc(x int, y int) (int, int) {
	return x + y, x - y
}

// 3.2 预先在参数返回
func calc1(x int, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	fmt.Println(sumFn(1, 2))

	sum, _ := calc1(2, 1)
	fmt.Println(sum)
}
