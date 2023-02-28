package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func addChan(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 1000)
		println(i, "成功插入i")
	}
	wg.Done()
}

func getChan(ch chan int) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Millisecond * 10)
		println(<-ch, "成功读取i")
	}
	wg.Done()
}

func main() {
	// 管道安全：管道在协程中是安全的
	// 虽然插入的时间间隔是1s,读取的时间间隔是10ms，但是 读取时还是会等待管道插入
	// 而不是在读取不到数据时直接报错

	var ch1 = make(chan int, 10)
	wg.Add(1)
	go getChan(ch1)
	wg.Add(1)
	go addChan(ch1)

	wg.Wait()

}
