package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	count := 0
	for t := range ticker.C {
		count++
		fmt.Println(t, count)
		if count > 10 {
			ticker.Stop()
		}
	}
}
