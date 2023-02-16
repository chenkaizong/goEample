package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1、 float32 float64
	var a float32 = 100.099
	var b float64 = 100

	fmt.Printf("num = %.2f,type is %T,size %d\n", a, a, unsafe.Sizeof(a))
	fmt.Printf("num = %f,type is %T,size %d\n", b, b, unsafe.Sizeof(b))

	//2、浮点数默认是float64
	c := 10.8
	fmt.Printf("num = %.2f,type is %T,size %d\n", c, c, unsafe.Sizeof(c))

	//3、科学计数法
	d := 10e3
	fmt.Printf("num = %.2f,type is %T,size %d\n", d, d, unsafe.Sizeof(d))

	//4、整型相除得到的是整数
	f := 8 / 10
	fmt.Printf("num = %d,type is %T,size %d\n", f, f, unsafe.Sizeof(f))

	// 5、精度丢失,解决办法  go get github.com/shopspring/decimal

	e := 0.1 + 0.2
	fmt.Printf("num = %f,type is %T,size %d\n", e, e, unsafe.Sizeof(e))

	// 6、转换
	g := float64(10)
	fmt.Printf("num = %f,type is %T,size %d\n", g, g, unsafe.Sizeof(g))
}
