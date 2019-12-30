package main

import "fmt"

func main() {

	// 数组定义： var name [n]T
	// n 为正整数。
	var a [0]int // 数组类型为[3]int
	var b [4]int // 数组类型为[4]int
	//a = b //不可以这样做，因为此时a和b是不同的类型

	fmt.Printf("%T   %T  \n",a,b)

	array := []int{1, 2, 3, 4, 5}
	temp := array
	fmt.Printf("%T %p \n", array, array)
	fmt.Printf("%T %p \n", temp, temp)
	slice := array[2:4]
	for i, v := range slice {
		fmt.Println(i, v)
	}
	slice = append(slice, 8)
	fmt.Printf("%p, %d, %d \n",slice, len(slice), cap(slice))
	slice = append(slice, 80)
	fmt.Printf("%p, %d, %d \n",slice, len(slice), cap(slice))
	fmt.Println(array)

	slice2 := array[:3]
	fmt.Println(slice2, len(slice2), cap(slice2))

	slice3 := slice2[:2]
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice3[0] = 100

	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(array)

	slice4 := array[:3]
	fmt.Printf("%p,%v \n",&slice4[1], slice4)
	slice4 = append(slice4[:1], slice4[2:]...)
	fmt.Printf("%p,%v \n",&slice4[1], slice4)
	fmt.Println(array)

	m := make(map[string]int, 2)
	fmt.Printf("%p %v \n", m, m)
	m["first"] = 1
	fmt.Printf("%p %v \n", m, m)
	m["second"] = 2
	m["third"] = 3
	fmt.Printf("%p %v \n", m, m)


}
