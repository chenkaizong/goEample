package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func run() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println(error)
		}
		wg.Done()
	}()

	panic("aaaaa")

}

func main() {
	wg.Add(1)
	go run()
	wg.Add(1)
	go run()

	wg.Wait()
}
