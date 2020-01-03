package main

import "fmt"

func main() {
	d1 := &Dog{
		Feet:   4,
		Animal: &Animal{name: "天天"},
	}

	d1.move()
	d1.wang()

	fmt.Printf("%p %p %p \n",&d1, &d1.Animal, &d1.name)

	d2 := d1
	fmt.Printf("%p %p %p \n",&d2, &d2.Animal, &d2.name)
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
	fmt.Printf("%s 旺旺 \n", d.name)
}