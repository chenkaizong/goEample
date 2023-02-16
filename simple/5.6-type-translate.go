package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1、常用数字互相转换：float64 int 等只要在后面加上括号即可

	// 2、转字符串
	a := 10
	fmt.Println(fmt.Sprintf("%d", a))
	b := 10.2
	fmt.Println(fmt.Sprintf("%f", b))
	c := '我'
	fmt.Println(fmt.Sprintf("%c", c))
	d := true
	fmt.Println(fmt.Sprintf("%t", d))

	//其中v%是最通用的
	fmt.Println(fmt.Sprintf("%v", c))

	// 3、专门的函数strconv
	// 3.1转字符串,FormatInt , FormatFloat, FormatBool, FormatByte等
	fmt.Println(strconv.FormatInt(int64(a), 10))
	fmt.Println(strconv.FormatFloat(float64(a), 'f', 2, 32))

	// 3.2 字符串转数字
	fmt.Println(strconv.ParseInt("3", 10, 64))
	fmt.Println(strconv.ParseFloat("3.14", 64))

}
