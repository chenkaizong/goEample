package main

import (
	"fmt"
	"sync"
)

var count = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func addCount() {
	mutex.Lock()
	count++
	fmt.Println(count)
	wg.Done()
	mutex.Unlock()
}
func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go addCount()
	}
	wg.Wait()

}
