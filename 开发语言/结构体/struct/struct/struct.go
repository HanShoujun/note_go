package main

import "fmt"

func main() {
	d1 := Dog{
		Feet:   4,
		Animal: &Animal{name: "天天"},
	}

	d1.move()
}

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s 动！\n", a.name)
}

type Dog struct {
	Feet int8
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s 旺旺", d.name)
}