package main

import "fmt"

func main() {

	//1、for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//2、for持续
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
	}
	//3、 for range

	var arr []string = []string{"a", "b", "c"}
	for index, str := range arr {
		fmt.Println(index, str)
	}
}
