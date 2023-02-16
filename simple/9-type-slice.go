package main

import "fmt"

func main() {
	// 1、切片
	var arr1 = [...]string{"hello", "h", "hhhh", "bbbb", "cccc", "dddd"}
	arr2 := arr1[1:3]
	fmt.Printf("%T\n", arr2)
	fmt.Println(arr2)

	//2、这是切片。切片可以使用append等切片方法
	arr3 := []string{"a", "b", 6: "c"}
	fmt.Println(arr3)
	arr4 := append(arr3, "aaa")
	fmt.Println(arr4)
	fmt.Printf("%v\n", arr4)

	//3、切片未赋值 是nil

	var arr6 []int
	fmt.Println(arr6 == nil)

	// 报错，不能比较
	// var arr7 [2]int
	// fmt.Println(arr7 == nil)

	//4、获取所有值
	arr8 := arr1[:]
	fmt.Println(arr8)

	// 5、长度和容量

	fmt.Println(len(arr4)) //8
	fmt.Println(cap(arr4)) // 14
}
