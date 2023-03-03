package main

import "fmt"

func main() {
	var ch1 = make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	var ch2 = make(chan string, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- fmt.Sprintf("%d", i)
		}
	}()

	// 监听的case中,没有满足条件的就阻塞
	// 多个满足条件的就任选一个执行
	// select本身不带循环,需要外层的for
	// default通常不用,会产生忙轮询
	// break只能跳出select中的一个case

	for {
		select {
		case v := <-ch1:
			fmt.Println("数字", v)
			break
		case v := <-ch2:
			fmt.Println("字符", v)
			break
		}
	}
}
