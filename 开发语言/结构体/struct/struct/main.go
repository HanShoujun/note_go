package main

import (
	"fmt"
	"struct/other"
)

func main() {

	var myInt MyInt = MyInt(12)
	var ivar = 123

	var aliasInt MyAlias = 345

	var runeChar rune = 'h'

	fmt.Println(myInt)
	fmt.Printf("%T \n", myInt) //main.MyInt
	fmt.Printf("%T \n", ivar) //int
	fmt.Printf("%T \n", aliasInt) //int

	fmt.Printf("%T \n", runeChar) //int32

	var p1 = person{
		name: "李三",
		city: "成都",
		age:  20,
	}
	var p2 person // {name:"", city:"", age:0}
	//p2.name = "王五"

	fmt.Printf("%T %#v \n",p1, p1)
	fmt.Printf("%T %#v \n",p2, p2)

	var p3 = new(person)
	p3.name = "赵六"
	fmt.Printf("%T %#v \n",p3, p3)

	mst := make(map[string]*student)
	students := []student{
		{
			name: "小王子",
			age:  18,
		},
		{
			name: "娜扎",
			age:  23,
		},
		{
			name: "大王八",
			age:  1000,
		},
	}

	for i, _ := range students {
		mst[students[i].name] = &students[i]
		fmt.Printf("%v %p \n",students[i], &students[i])
	}
	for k, v := range mst {
		fmt.Println(k, "=>", v)
	}

	p5 := person{
		name: "张三",
		city: "大连",
		age:  33,
	}
	p5.dream()

	p6 := NewPerson("王二麻子", "幸福村", 19)
	p6.dream()

	// p5 age：33
	p5.addAge(3)
	fmt.Println(p5.name,p5.age)
	p5.subAge(6)
	fmt.Println(p5.name,p5.age)

	// 其他包的结构体
	customStruct := other.CustomStruct{
		Name: "hi",
		Des:  "description",
	}
	fmt.Println(customStruct)
	customStruct.ChangeName("HELLO")
	fmt.Println(customStruct)

	user := User{
		Name:    "小豆子",
		Gender:  "男",
		Address: Address{
			Province:	"山东",
			City:		"青岛",
		},
	}
	fmt.Printf("%#v \n", user)

	teacher := Teacher{
		Name:  "李老师",
		Class: "五年级二班",
		Address: Address{
			Province: "河北",
			City:     "石家庄",
		},
	}
	fmt.Println(teacher.Address.Province)
	fmt.Println(teacher.City)
	fmt.Println(teacher.Province)
}

type MyInt int

type MyAlias = int

type person struct {
	name string
	city string
	age int8
}

type student struct {
	name string
	age int
}

func NewPerson(name string, city string, age int8) *person {
	return &person{
		name:name,
		city:city,
		age:age,
	}
}

func (p person) dream()  {
	fmt.Printf("%s dream \n", p.name)
}

// 值类型的接收者，在方法中对p做修改，不会影响原变量
func (p person) subAge(age int8) {
	p.age -= age
}

// 引用类型的接收者，在方法中对p做修改，会影响原变量
func (p *person) addAge(age int8) {
	p.age += age
}

// error: cannot define new methods on non-local type other.CustomStruct
//func (c other.CustomStruct) ChangeDes(des string) {
//	fmt.Println(des)
//}

// 结构体嵌套
type Address struct {
	Province string
	City string
}

type User struct {
	Name string
	Gender string
	Address Address
}

// 嵌套匿名结构体
type Teacher struct {
	Name string
	Class string
	Address
}

