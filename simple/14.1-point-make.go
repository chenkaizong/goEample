package main

import "fmt"

// make和new都是golang用来分配内存的內建函数，且在堆上分配内存，make 即分配内存，也初始化内存。new只是将内存清零，并没有初始化内存。
// make返回的还是引用类型本身；而new返回的是指向类型的指针。
// make只能用来分配及初始化类型为slice，map，channel的数据；new可以分配任意类型的数据。

func main() {
	//make 就是创建指针类型，并且分配存储空间，只能用于 map
	arr1 := []int{1, 2, 3}
	arr2 := arr1
	arr2[2] = 100
	fmt.Println(arr1, arr2)

	var userinfo = make(map[string]string, 3)
	userinfo["HAHA"] = "我的"
	fmt.Println(userinfo)

	var arr = make([]int, 3, 6)
	arr[1] = 100
	fmt.Println(arr)

}
