package main

import (
	"fmt"
	"time"
)

func run() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}

// 1、使用go 即可实现并行
// 2、主进程执行完毕后 协程也会退出

func main() {
	go run()
	go run()
	// 接受命令行输入, 不做任何事情，相当于主进程都不退出了
	var input string
	fmt.Scanln(&input)

}
