package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var flag = true
	fmt.Printf("type is %T,size %d\n", flag, unsafe.Sizeof(flag))

}
