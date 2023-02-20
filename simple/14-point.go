package main

import "fmt"

func main() {
	a := 10
	b := &a
	*b = 11
	fmt.Println(b, a)

	arr1 := []int{1, 2, 3}
	arr2 := arr1
	arr2[2] = 100
	fmt.Println(arr1, arr2)

}
