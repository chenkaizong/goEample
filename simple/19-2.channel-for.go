package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var wg sync.WaitGroup

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

func printPrime(start int, end int) {
	for i := start; i <= end; i++ {
		if IsPrime(i) {
			println(i, "是素数")
		}
	}
	wg.Done()
}

func main() {
	// var ch1 = make(chan int, 4)
	starttime := time.Now().UnixMilli()

	// 1、直接算
	// printPrime(1, 10000)  //452ms

	// 2、协程序 //425ms
	var ch1 = make(chan int, 4)
	for i := 0; i < 4; i++ {
		ch1 <- i
	}
	close(ch1)
	for v := range ch1 {
		wg.Add(1)
		fmt.Println(v)
		go printPrime(v*2500+1, (v+1)*2500)
	}

	wg.Wait()
	endtime := time.Now().UnixMilli()
	println("开始时间", starttime, "总用时:", endtime-starttime)

}
