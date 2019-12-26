package main

import "fmt"

func modifyArray(x [3]int) {
	fmt.Printf("modifyArray:%p \n",&x)
	x[0] = 100
	fmt.Printf("modifyArray:%p \n",&x)
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
func main() {

	s := "heel 是大佛"
	for _, c := range s {
		fmt.Printf("%c ",c)
	}

	fmt.Println("*************")

	a := [3]int{10, 20, 30}
	a1 := a
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Printf("%p \n",&a)
	fmt.Printf("%p \n",&a1)
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

func printAddr(s string) {
	var temp string = s
	fmt.Printf("%p \n", &s)
	fmt.Printf("%p \n", &temp)
}