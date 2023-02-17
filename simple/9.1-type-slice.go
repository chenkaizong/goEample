package main

import "fmt"

func main() {
	// 1切片的函数
	slice := make([]int, 4, 10)
	fmt.Println(slice)
	//2 新增
	slice = append(slice, 3)
	fmt.Println(slice)
	// 3 复制，因为是引用类型，所以可以用copy来复制
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	// 4、倒叙
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	// 5、合并

	slice = append(slice, sliceCopy...)

	fmt.Println(slice)
	// 6、切片是没有直接的删除函数，只能通过 append 实现
}
