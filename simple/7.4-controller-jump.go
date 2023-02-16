package main

import "fmt"

func main() {
	// 跳出循环
	//1、continue 跳出本次,

	//2、break 跳出本循环

	//3、在多次循环的时候，可以通过标签label可以跳出多个循环，label必须紧跟循环体
MY_LABEL:
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if j%2 == 0 {
				continue MY_LABEL
			}
			fmt.Printf("i=%d,j=%d\n", i, j)
		}
	}

	//4、goto 可以在控制器里使用 GOTO_TAG
	a := 10
	if a == 10 {
		goto GOTO_TAG
	}

	fmt.Println(a)
GOTO_TAG:
	fmt.Println("goto")

}
