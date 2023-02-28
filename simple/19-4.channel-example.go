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

func printPrimeChan(ch chan int, i int) {
	for v := range ch {
		if IsPrime(v) {
			// println(v, "是素数(协程号", i, ")")
		}
	}
	wg.Done()
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
	wg.Add(1)
	go addNumChan(ch1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printPrimeChan(ch1, i)
	}

	wg.Wait()
	endtime := time.Now().UnixMilli()
	fmt.Println("耗时", endtime-starttime)
}
