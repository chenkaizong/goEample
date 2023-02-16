package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// int
	var a int = 100
	fmt.Println(a)
	//1.1 数字类型  int8 int16 int32 int64
	//1.2 数字类型无符号 uint8 uint16 uint32 uint64
	//1.3 数字类型  int: 在 64 位处理器上占 64 位，在 32 位处理器上面占 32 位，所以最好用 int64
	//    特别注意，牵扯到文件传输，数据交换等牵扯掉不同系统的 一定不能使用int

	//2、占位 2 int8 占用1位字节，int16 2位  int32 4位 int64 8位
	fmt.Println(unsafe.Sizeof(a)) // 8

	//3、数字类型转换，不同的数字类型是不能进行运算的，所以必须知道转换方式

	var b int16 = 1000
	var c int32 = 10000
	// fmt.Println(b + c) 报错
	fmt.Println(int32(b) + c)

	d := 100
	fmt.Printf("num = %d,type is %T,size %d", d, d, unsafe.Sizeof(d))
}
