package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func run(i int) {
	var times int
	for times < 10 {
		times++
		fmt.Printf("协程%v,第%v次\n", i, times)
	}
	defer wait.Done()
}

// 1、使用go 即可实现并行
// 2、主进程执行完毕后 协程也会退出

func main() {
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go run(i)
	}
	wait.Wait()

}
