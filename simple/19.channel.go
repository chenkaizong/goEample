package main

func main() {
	// 声明一个 int 的管道类型,并分配空间
	ch1 := make(chan int, 3)

	// 存储管道
	ch1 <- 10
	ch1 <- 11
	ch1 <- 12
	// ch1 <- 13  死锁，不能超过容量

	// 获取管道里面的内容,先进先出
	a := <-ch1
	println(a)
	b := <-ch1
	println(b)
	c := <-ch1
	println(c)
	// d := <-ch1
	// println(d) 全部取完后，再取报错

	ch1 <- 12
	f := <-ch1
	println(f) //
}
