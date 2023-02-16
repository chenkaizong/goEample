package main

import "fmt"

func main() {
	// 1、先声明后赋值
	var arr [3]int
	arr = [3]int{1, 2, 4}
	arr[1] = 100
	fmt.Println(arr)
	fmt.Printf("%T\n", arr) //[3]int ,注意，当有定义长度的时候，数组长度也是类型的一部分

	// 2、先声明时赋值, []不写长度也是为3
	var arr2 = []int{1, 2, 3}
	fmt.Printf("%T\n", arr2)
	// arr2 = arr //虽然长度为3，但不能赋值[3]int 是不能赋值的

	// 3、自行判断长度
	var arr3 = [...]int{1, 2, 3}
	fmt.Printf("%T\n", arr3) //[3]int ,
	arr3 = arr               //可以赋值

	// 4、短标签方式
	arr4 := [...]int{1, 2, 3}
	fmt.Printf("%T\n", arr4) //[]int ,

	// 5、骚操作
	arr5 := [...]int{1: 1, 2: 3, 100: 3}
	fmt.Printf("%T\n", arr5) //[101]int

	// 使用len 判断长度， rang可以遍历
	fmt.Println(len(arr2))

	// 6、多维数组

	arr10 := [...][3]int{ //只有第一层可以用...
		{1, 2, 3},
		{2, 3, 4},
		{4, 5, 6},
	}

	fmt.Println(arr10)
}
