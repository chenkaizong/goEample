package main

import "fmt"

func main() {
	a := 10
	b := &a
	*b = 11
	fmt.Println(b, a)

	fmt.Printf("%d,%T,%v\n", a, a, &a)   //11,int,0xc00001a088
	fmt.Printf("0x%x,%T,%v\n", b, b, &b) //0xc00001a088,*int,0xc00000a028

}
