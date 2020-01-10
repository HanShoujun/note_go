package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 3.14
	b := 100
	reflectType(a)
	reflectType(b)

	var p *float32
	var m myInt
	var c rune
	reflectType(p)
	reflectType(m)
	reflectType(c)

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "沙河小王子",
		age:  18,
	}
	var e = book{title: "《跟小王子学Go语言》"}

	reflectType(d)
	reflectType(e)
}

type myInt int64

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v \n",v.Name(), v.Kind())
}
