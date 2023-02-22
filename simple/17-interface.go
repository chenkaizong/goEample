package main

import "fmt"

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (phone Phone) start() {
	fmt.Println("start")
}

func (phone Phone) stop() {
	fmt.Println("stop")
}

func main() {
	var p1 Usber
	p1 = Phone{}
	p1.start()

}
