package main

import (
	"fmt"

	"example.com/m/calc"
)

// 1、 go mod init example.com/m
// 2、
func main() {
	// fmt.Println(calc.aaa) 这里肯定是报错的，不能是小写的
	fmt.Println(calc.Add(16, 20))
}
