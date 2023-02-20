package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	for count < 10 {
		count++
		time.Sleep(time.Second)
		fmt.Println("延迟执行")
	}
}
