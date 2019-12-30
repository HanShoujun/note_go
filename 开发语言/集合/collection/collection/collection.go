package main

import "fmt"

func main() {

	// 数组定义： var name [n]T
	// n 为常量。
	var a [3]int // 数组类型为[3]int
	var b [4]int // 数组类型为[4]int
	//a = b //不可以这样做，因为此时a和b是不同的类型

	fmt.Printf("%T:%v  \n",a,a)
	fmt.Printf("%T:%v  \n",b,b)

	// 如果有初始值，可以用 [...] 。让编译器推断出类型
	var c = [...]int{2,3,4}
	fmt.Printf("%T:%v  \n",c,c)

	for _, v := range c {
		fmt.Println(v)
	}

	// 切片定义： var name []T  （与数组非常相近）
	var sliceA []int
	fmt.Printf("%T %v %t \n", sliceA, sliceA, sliceA == nil)

	// 切片是对数组的封装，切片的操作其实是对底层数组的操作
	array := []int{1, 2, 3, 4, 5}
	temp := array
	fmt.Printf("%T %p %p \n", array, array, &array[2])
	fmt.Printf("%T %p \n", temp, temp)
	slice := array[2:3]
	for i, v := range slice {
		fmt.Println(i, v)
	}
	slice = append(slice, 8)
	fmt.Printf("%p, %d, %d \n",slice, len(slice), cap(slice))
	slice = append(slice, 80)
	fmt.Printf("%p, %d, %d \n",slice, len(slice), cap(slice))
	fmt.Println(array)
	// 当底层数组容量不够时，会创建一个新的数组
	slice = append(slice, 800)
	// 再次操作切片将只会影响新数组，不会影响原来的数组
	slice[0] = 999
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

	var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	// 所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。
	fmt.Println(len(s1), len(s2), len(s3))

	delSlice()

	copySlice()

	// make(map[KeyType]ValueType, [cap])
	// 其中cap表示map的容量，该参数虽然不是必须的，
	// 但是我们应该在初始化map的时候就为其指定一个合适的容量。
	m := make(map[string]int, 2)
	fmt.Printf("%p %v %d \n", m, m, len(m))
	m["first"] = 1
	fmt.Printf("%p %v %d \n", m, m, len(m))
	m["second"] = 2
	m["third"] = 3
	fmt.Printf("%p %v %d \n", m, m, len(m))

	// 判断某个键是否存在
	// value, ok := map[key]
	value, ok := m["second"]
	fmt.Println(value, ok)
	value2, ok := m["fourth"]
	fmt.Println(value2, ok)

	for key, v := range m {
		fmt.Println(key,v)
	}

	delete(m, "second")
	fmt.Println(m)
}

func copySlice() {
	intSlice := []int{1, 2, 3, 4, 5}
	newSlice := make([]int, 3, 4)
	copy(newSlice, intSlice)
	fmt.Println(intSlice)
	fmt.Println(newSlice)

	newSlice[0] = 90
	fmt.Println(intSlice)
	fmt.Println(newSlice)

	fmt.Println("*************")
}

func delSlice() {
	intSlice := []int{1, 2, 3, 4, 5}

	slice4 := intSlice[:3]
	fmt.Printf("%p,%v \n",&slice4[1], slice4)
	slice4 = append(slice4[:1], slice4[2:]...)
	fmt.Printf("%p,%v \n",&slice4[1], slice4)
	fmt.Println(intSlice)

	fmt.Println("*************")
}