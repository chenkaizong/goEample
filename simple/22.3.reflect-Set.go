package main

import (
	"fmt"
	"reflect"
)

func reflectSet(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Elem().Kind() // 当传进来的是指针时，v.Kind()是ptr
	switch k {
	case reflect.Int:
		v.Elem().SetInt(111)
	case reflect.Float32:
		v.Elem().SetFloat(111.1)
	}
}

func main() {
	var a int = 100
	var b float32 = 100.1
	reflectSet(&a)
	reflectSet(&b)

	fmt.Println(a, b)

}
