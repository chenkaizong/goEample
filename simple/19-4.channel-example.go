package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 加入所有要判断的数
func addNumChan(ch chan int) {
	for i := 2; i < 10000; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func PrimeChan(ch chan int, printCh chan int, runCh chan bool) {
	for v := range ch {
		if IsPrime(v) {
			printCh <- v
		}
	}
	runCh <- true
	if len(runCh) == cap(runCh) {
		close(printCh)
	}
	wg.Done()
}

func printChan(printCh chan int) {
	for v := range printCh {
		println(v, "是素数")
	}
}

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	//偶数一定不是素数
	if n > 2 && n%2 == 0 {
		return false
	}
	//从2遍历到n的方根，看看是否有因子
	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			//发现一个因子
			return false
		}
	}
	return true
}

func main() {
	starttime := time.Now().UnixMilli()
	var ch1 = make(chan int, 10000)
	var printCh = make(chan int, 100)
	var runCount = make(chan bool, 10)
	wg.Add(1)
	go addNumChan(ch1)
	for i := 0; i < len(runCount); i++ {
		wg.Add(1)
		go PrimeChan(ch1, printCh, runCount)
	}
	wg.Add(1)
	go printChan(printCh)

	wg.Wait()
	endtime := time.Now().UnixMilli()
	fmt.Println("耗时", endtime-starttime)
}
