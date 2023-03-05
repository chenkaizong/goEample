package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name  string
	Age   int
	Score float64
}

func PrintStructFields(s interface{}) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("不是结构题类")
	}
}

func main() {
	var person Person = Person{"name", 10, 11.2}
	PrintStructFields(person)
}
