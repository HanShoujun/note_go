package main

import "fmt"

func main() {

	var x Sayer
	c := cat{}
	d := dog{}
	//x.say() // panic nil pointer error
	x = c // 可以接收结构体类型
	x.say()
	x = &d // 可以接收指针类型
	x.say()

	f := fish{}
	//x = f // 不可以接收结构体类型
	x = &f // 可以接收指针类型
	f.say()

	d.move()

	t := toyDog{}
	x = t
	x.say()

	var m Mover = t
	m.move()

	var ani animal = t
	ani.move()

	// 空接口可以存储任意类型
	var all interface{} = ani

	showInfo(all)
	showInfo(ani)

	// 类型断言
	d2,ok := all.(Sayer)
	if ok {
		d2.say()
	}

	fmt.Println("********")
	justify(ani)
	justify(m)
	justify(&f)
}

type Sayer interface {
	say()
}

type dog struct {
	name string
}

type cat struct {

}

// 值接收者实现接口
func (d dog) say() {
	fmt.Println("旺旺")
}

func (c cat) say() {
	fmt.Println("喵喵")
}

type fish struct {

}

// 指针接收者实现接口
func (f *fish) say() {
	fmt.Println("无声")
}

type Mover interface {
	move()
}

// 实现多个接口
func (d dog) move() {
	fmt.Println(d.name, "会动")
}

// 实现接口可以"继承"
type toyDog struct {
	dog
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

func showInfo(info interface{}) {
	fmt.Printf("%T %v \n", info, info)
}

// 类型判断
func justify(x interface{}) {
	switch v := x.(type) {
	case Sayer:
		v.say()
	case Mover:
		v.move()
	case animal:
		v.move()
		v.say()
	default:
		fmt.Println("unknown")
	}
}