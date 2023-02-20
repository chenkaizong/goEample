package main

import (
	"errors"
	"fmt"
)

// 使用 panic 和 recover 来 抛出和捕获错误， 在这里必须使用 defer配合
func fn1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("抛出错误")
}

func fn2(a, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	return a / b
}

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("文件读取失败")
	}
}

func myFn() {
	defer func() {
		error := recover()
		if error != nil {
			fmt.Println("给管理员发送游戏键")

		}
	}()
	err := readFile("xxx.go")
	if err != nil {
		panic(err)
	}
}

func main() {
	fn1()
	res := fn2(10, 2)
	fmt.Println(res)
	fmt.Println(fn2(5, 0))

	myFn()
}
