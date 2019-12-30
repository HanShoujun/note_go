package main

import "fmt"

func ifDemo() {
	score := 65
	if score > 95 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func forDemo() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func switchDemo() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效")
	}

}

func switchDemo2(age int) {

	switch {
	case age < 25:
		fmt.Println("好好学习")
	case age > 25 && age < 35:
		fmt.Println("好好工作")
	case age > 60:
		fmt.Println("好好活着")
	}
}

func main() {

	ifDemo()

	forDemo()

	switchDemo()

	switchDemo2(33)

	fmt.Println("*************")

	s := "heel 是大佛"
	for _, c := range s {
		fmt.Printf("%c ", c)
	}

	fmt.Println("*************")

	a := [3]int{10, 20, 30}
	a1 := a
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Printf("%p \n", &a)
	fmt.Printf("%p \n", &a1)
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]

	s2 := "a string"
	s3 := s2
	fmt.Printf("%p \n", &s2)
	fmt.Printf("%p \n", &s3)
	printAddr(s2)

	intSlice := []int{2, 3}
	others := intSlice
	others[0] = 9
	fmt.Printf("intSlice :%p,%v \n", intSlice, intSlice)
	fmt.Printf("others :%p,%v \n", others, others)
}

func modifyArray(x [3]int) {
	fmt.Printf("modifyArray:%p \n", &x)
	x[0] = 100
	fmt.Printf("modifyArray:%p \n", &x)
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func printAddr(s string) {
	var temp string = s
	fmt.Printf("%p \n", &s)
	fmt.Printf("%p \n", &temp)
}
