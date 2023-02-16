package main

import "fmt"

// 一个
const pi = 3.1415926

// 多个,未定义的默认与 上面一致
const (
	N1 = 3
	N2 = "hahah"
	N3 //hahah
)

// iota 计数器， 第一个为0 接下来为1，2，3
const (
	M1 = iota       //0
	M2 = 100        //100
	M3              //100
	M4 = iota       //3
	_               //忽略
	M5              // 4
	M6 = iota * 100 // 500
)

func main() {
	fmt.Println(pi)
	fmt.Println(N1, N2, N3)
	fmt.Println(M1, M2, M3, M4, M5, M6)
}
