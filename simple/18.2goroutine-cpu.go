package main

import (
	"fmt"
	"runtime"
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
	// 一般多核 cpu, 设置 cpu 数目减一的效率要比直接设置 runtime.NumCPU () 高一点。
	maxProces := runtime.NumCPU()
	if maxProces > 1 {
		maxProces -= 1
	}
	runtime.GOMAXPROCS(maxProces)

	go run()
	go run()
	// 接受命令行输入, 不做任何事情，相当于主进程都不退出了
	var input string
	fmt.Scanln(&input)

}
