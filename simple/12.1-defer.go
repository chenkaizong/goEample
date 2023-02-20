package main

import "fmt"

// defer
// 1、defer 会Go语言的 defer 语句会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，
// 让函数或语句可以在当前函数执行完毕后（包括通过return正常结束或者panic导致的异常结束）执行
// 2、将延迟处理的语句按 defer 的逆序进行执行，

func f1() {
	fmt.Println("开始")
	defer fmt.Println("1111")
	defer fmt.Println("2222")
	fmt.Println("结束")
}

// 这个函数的返回值是6 而不是5
func f2() (a int) {
	defer func() {
		a++
	}()
	return 5
}

func f3(a int) int {
	defer func() {
		a++
	}()
	return a
}

func main() {
	f1()
	// 输出：
	// 开始
	// 结束
	// 2222   //defer是逆序的
	// 1111
	fmt.Println(f2()) //输出6，   defer是在return 之前输出的
	fmt.Println(f3(0))
}
