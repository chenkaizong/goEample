package main

import (
	"fmt"
	"sort"
)

func main() {
	// 数组和切片是可以排序的
	arr1 := [...]int{1, 6, 3, 7, 5}
	arr2 := [...]string{"hah", "bbb", "ccc", "aaa", "mmm"}
	fmt.Println(arr1)
	fmt.Println(arr2)
	sort.Ints(arr1[:])
	sort.Strings(arr2[:])
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Printf("%v,%T", arr1, arr1)

	arr3 := []int{1, 6, 3, 7, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(arr3)))
	fmt.Println(arr3)

	s := []int{5, 2, 6, 3, 1, 4} // 未排序的切片数据
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s) //将会输出[6 5 4 3 2 1]

}
