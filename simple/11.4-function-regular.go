package main

import "fmt"

// 1、递归调用
func sum(x int) int {
	if x == 1 {
		return x
	} else {
		return x + sum(x-1)
	}
}

// 2、闭包

func close() func(int) int {
	res := 0
	return func(y int) int {
		res += y
		return res
	}
}

func main() {
	fmt.Println(sum(100))
	closeFunc := close()
	fmt.Println(closeFunc(10))
	fmt.Println(closeFunc(11))
}
