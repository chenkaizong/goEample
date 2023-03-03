package main

import (
	"fmt"
)

func main() {
	var intChan = make(chan int, 10)
	var stringChan = make(chan string, 10)
	go func() {
		for i := 0; i < 10; i++ {
			intChan <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			stringChan <- fmt.Sprintf("%d字符串", i)
		}
	}()

	// 使用select 的时候不得使用 close 来关闭管道

	// 监听的case中,没有满足条件的就阻塞
	// 多个满足条件的就任选一个执行
	// select本身不带循环,需要外层的for
	// default通常不用,会产生忙轮询
	// break只能跳出select中的一个case

	for {
		select {
		case v := <-intChan:
			fmt.Println("数字", v)
			break
		case v := <-stringChan:
			fmt.Println("字符串", v)
			break
		}
	}
}
