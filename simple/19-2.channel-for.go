package main

import (
	"fmt"
	"math"
	"time"
)

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
}

func main() {
	// var ch1 = make(chan int, 4)
	starttime := time.Now().UnixMilli()

	// 1、直接算
	// printPrime(1, 1000)

	// 2、协程序
	var ch1 = make(chan int, 4)
	for i := 0; i < 4; i++ {
		ch1 <- i
	}
	close(ch1)
	for v := range ch1 {
		fmt.Println(v)
		go printPrime(v*250+1, (v+1)*250)
	}

	endtime := time.Now().UnixMilli()
	println("开始时间", starttime, "总用时:", endtime-starttime)

}
