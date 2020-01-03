package other

import (
	"fmt"
	"math/rand"
	"package/tool"
)

func init() {
	fmt.Println("other package init")
}

// 小写字母不可导出
func random(n int) int {
	network := tool.Network{
		Name:"www.baidu.com",

	}
	network.FindByName("Baidu")
	return rand.Intn(n)
}

// 大写字母开头，可导出
func Add(x int, y int) int {
	return x+y
}

func WrapRandom(n int) int {
	return random(n)
}