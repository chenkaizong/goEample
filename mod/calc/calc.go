package calc

import "fmt"

var aaa int = 100

// 小写的表示私有
func add(x, y int64) int64 {
	return x + y
}

// 大写的表示公有的可以调用
func Add(x, y int64) int64 {
	fmt.Println(aaa)
	return x + y
}
