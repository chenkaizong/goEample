package main

import (
	"fmt"
	"sync"
	"time"
)

func run() {
	var times int
	for times < 10 {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}

// 1、使用go 即可实现并行
// 2、主进程执行完毕后 协程也会退出

func main() {
	var wait sync.WaitGroup
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func() {
			run()
			defer wait.Done()
		}()
	}
	wait.Wait()

}
