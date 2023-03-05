package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name  string `json:"name" form:"fieldname"`
	Age   int
	Score float64
}

func (Person) Method1() {

}

func (Person) Aaa() {

}

func PrintStructFields(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("不是结构题类")
		return
	}

	// field0 := t.Field(0) //索引方式
	field0, _ := t.FieldByName("Name") //名称索引

	fmt.Println(field0) //{Name  string  0 [0] false}
	fmt.Printf("名称：%v，类型:%v\n", field0.Name, field0.Type)
	fmt.Printf("字段tag:%v\n", field0.Tag.Get("json"))
	fmt.Printf("字段tag:%v\n", field0.Tag.Get("form"))

	var fieldCount = t.NumField()
	fmt.Printf("结构体属性个数%v\n", fieldCount)

	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
	fmt.Println(v.FieldByName("Score"))

	// 遍历

	for i := 0; i < fieldCount; i++ {
		fmt.Printf("属性名称：%v 属性值%v，属性类型%v, 属性Tag%v\n", t.Field(i).Name, v.Field(i), t.Field(i).Type, t.Field(i).Tag.Get("json"))
	}

}

func PrintStructFn(s interface{}) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("不是结构题类")
		return
	}

	m := t.Method(0) //顺序是按照方法名的 assi码来操作的

	fmt.Printf("%v\n", m.Name)
}

func main() {
	var person Person = Person{"name", 10, 11.2}
	// PrintStructFields(person)
	PrintStructFn(person)

}
