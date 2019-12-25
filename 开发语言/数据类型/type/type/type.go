package main

import (
	"fmt"
	"strings"
)

func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)  // 10
	fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b)  // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)  // ff
	fmt.Printf("%X \n", c)  // FF

	fmt.Println("********************")
	var n = 100
	fmt.Printf("%T \n", n)
	fmt.Printf("%v \n", n)
	fmt.Printf("%b \n", n)
	fmt.Printf("%d \n", n)
	fmt.Printf("%o \n", n)
	fmt.Printf("%x \n", n)

	var s = "hei 胜多负少 322434"
	fmt.Printf("%T \n", s)
	fmt.Printf("%s \n", s)
	fmt.Printf("%#v \n", s)

	fmt.Println("********************")
	s2 := "hello 你好"
	s3 := `第一个
	第二个sdf
	第三个234`
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(s2 + s3)
	split := strings.Split(s3, "\n")
	fmt.Println(split)
	fmt.Println(strings.Contains(s2, "llo"))
	fmt.Println(strings.HasPrefix(s3, "第一个"))
	fmt.Println(strings.LastIndex(s2, "l"))
	fmt.Println(strings.Join(split, "+"))

	fmt.Println("********************")
	c1 := '中'
	c2 := 'x'
	fmt.Printf("%T \n", c1)
	fmt.Printf("%T \n", c2)

	s4 := "234helf是大佛"
	for i := 0; i < len(s4); i++ {
		fmt.Printf("%v(%c) ", s4[i], s4[i])
	}
	fmt.Println()

	for _, r := range s4{
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
