package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func addChan(ch chan<- int) {
	for i := 0; i < 20; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 1000)
		println(i, "成功插入i")
	}
	wg.Done()
}

func getChan(ch <-chan int) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Millisecond * 10)
		println(<-ch, "成功读取i")
	}
	wg.Done()
}

func main() {
	var ch1 chan<- int = make(chan<- int, 10)
	ch1 <- 10
	// <-ch1  报错

	var ch2 <-chan int = make(<-chan int, 20)
	// ch2 <- 10 报错
	<-ch2

	var ch3 = make(chan int, 10)
	wg.Add(1)
	go getChan(ch3)
	wg.Add(1)
	go addChan(ch3)

	wg.Wait()

}
